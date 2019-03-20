package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/mrbenshef/SmartHomeAdapters/microservice/robotserver"
	"github.com/mrbenshef/SmartHomeAdapters/microservice/thermostatserver"

	"github.com/ory/dockertest"

	"google.golang.org/grpc/test/bufconn"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)
import empty "github.com/golang/protobuf/ptypes/empty"

var lis *bufconn.Listener
var servoRequests = []*robotserver.ServoRequest{}

func TestMain(m *testing.M) {
	// connect to docker
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// start thermodb
	resource, err := pool.Run("smarthomeadapters/thermodb", "latest", []string{"POSTGRES_PASSWORD=password"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	url = fmt.Sprintf("localhost:%s", resource.GetPort("5432/tcp"))

	// wait till db is up
	if err = pool.Retry(func() error {
		var err error
		db, err := sql.Open("postgres", fmt.Sprintf("postgres://postgres:password@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), database))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// start test gRPC server
	lis = bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()

	thermostatserver.RegisterThermostatServerServer(s, &server{
		DB:          getDb(),
		RobotClient: mockRobotServerClient{},
	})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	exitCode := m.Run()

	pool.Purge(resource)

	os.Exit(exitCode)
}

type mockRobotServerClient struct {
	thisIsAMock bool
}

func (c mockRobotServerClient) SetServo(ctx context.Context, in *robotserver.ServoRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	servoRequests = append(servoRequests, in)
	return &empty.Empty{}, nil
}

func (c mockRobotServerClient) SetLED(ctx context.Context, in *robotserver.LEDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("Call to SetLED!")
}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestGetThermostat(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := thermostatserver.NewThermostatServerClient(conn)

	thermostat, err := client.GetThermostat(context.Background(), &thermostatserver.ThermostatQuery{
		Id: "qwerty",
	})

	if err != nil {
		t.Errorf("Expected nil error, got: %v", err)
	}

	expectedThermostat := &thermostatserver.Thermostat{
		Id:            "qwerty",
		Tempreture:    293,
		MinAngle:      30,
		MaxAngle:      170,
		MinTempreture: 283,
		MaxTempreture: 303,
		IsCalibrated:  true,
	}

	if !reflect.DeepEqual(thermostat, expectedThermostat) {
		t.Errorf("Robots differ. Expected: %+v, Got: %+v", expectedThermostat, thermostat)
	}
}

func TestSetThermostat(t *testing.T) {
	cases := []struct {
		thermostatRequest    *thermostatserver.SetThermostatRequest
		expectedServoRequest *robotserver.ServoRequest
	}{
		{
			thermostatRequest: &thermostatserver.SetThermostatRequest{
				Id:         "qwerty",
				Tempreture: 283,
				Unit:       "kelvin",
			},
			expectedServoRequest: &robotserver.ServoRequest{
				RobotId: "qwerty",
				Angle:   30,
			},
		},
		{
			thermostatRequest: &thermostatserver.SetThermostatRequest{
				Id:         "qwerty",
				Tempreture: 293,
				Unit:       "kelvin",
			},
			expectedServoRequest: &robotserver.ServoRequest{
				RobotId: "qwerty",
				Angle:   100,
			},
		},
		{
			thermostatRequest: &thermostatserver.SetThermostatRequest{
				Id:         "qwerty",
				Tempreture: 303,
				Unit:       "kelvin",
			},
			expectedServoRequest: &robotserver.ServoRequest{
				RobotId: "qwerty",
				Angle:   170,
			},
		},
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := thermostatserver.NewThermostatServerClient(conn)

	for _, c := range cases {
		servoRequests = []*robotserver.ServoRequest{}

		stream, err := client.SetThermostat(ctx, c.thermostatRequest)

		for {
			status, err := stream.Recv()
			if err != nil {
				t.Errorf("Error receiving: %v", err)
			}

			if status.Status == thermostatserver.SetThermostatStatus_DONE {
				break
			}
		}

		if err != nil {
			t.Errorf("Expected nil error, got: %v", err)
		}

		if len(servoRequests) != 1 {
			t.Errorf("Expected single servo request, got: %+v (len = %d)", servoRequests, len(servoRequests))
		}

		if !reflect.DeepEqual(servoRequests[0], c.expectedServoRequest) {
			t.Errorf("Expected request: %+v, got %+v", c.expectedServoRequest, servoRequests[0])
		}
	}
}