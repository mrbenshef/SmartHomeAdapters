package com.github.halspals.smarthomeadapters.smarthomeadapters


import android.content.Context
import android.os.Bundle
import android.support.v4.app.Fragment
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.*
import com.github.halspals.smarthomeadapters.smarthomeadapters.model.Robot
import kotlinx.android.synthetic.main.activity_authentication.*
import okhttp3.ResponseBody
import org.jetbrains.anko.design.snackbar
import org.jetbrains.anko.toast
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class RobotFragment : Fragment() {

    private val fTag = "RobotFragment"

    private lateinit var robotId: String
    private var robot: Robot? = null

    private lateinit var progressBar: ProgressBar
    private lateinit var switch: Switch
    private lateinit var seekBar: SeekBar

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {

        // get robotId from bundle
        val robotIdArgument = arguments?.getString("robotId")
        if (robotIdArgument == null) {
            // no id passed, try to go back
            Log.d(tag, "No robotId passed to robotFragment")
            context?.toast("Oops, something went wrong")
            fragmentManager?.popBackStack()
            return null
        }
        robotId = robotIdArgument

        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_robot, container, false)
    }


    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        progressBar = view.findViewById(R.id.progress_bar)
        switch = view.findViewById(R.id.robot_switch)
        seekBar = view.findViewById(R.id.robot_seek_bar)

        // set initial visibility
        progressBar.visibility = View.VISIBLE
        switch.visibility = View.INVISIBLE
        seekBar.visibility = View.INVISIBLE

        fetchRobot()
    }


    /**
     * Fetches the robot with id of [robotId] and calls [onReceiveRobot]
     */
    private fun fetchRobot() {
        // TODO: fetch robot from server
        // TODO: remove test code

        (activity as MainActivity).restApiService.getRobot(robotId).enqueue(object: Callback<Robot> {
            override fun onFailure(call: Call<Robot>, t: Throwable) {
                Log.e(fTag, t.message)
            }

            override fun onResponse(call: Call<Robot>, response: Response<Robot>) {
                Log.d(fTag, "received")
                Log.d(fTag, "body: ${response.body()}")
                onReceiveRobot(response.body()!!)
            }
        })
    }

    /**
     * Called whenever a new [Robot] is received
     *
     * @param robot the robot
     */
    private fun onReceiveRobot(robot: Robot) {
        Log.d(fTag, "receive propagating")
        this.robot = robot

        Log.d(fTag, robot.toString())

        progressBar.visibility = View.INVISIBLE
        switch.visibility = View.INVISIBLE
        seekBar.visibility = View.INVISIBLE

        when(robot.robotInterfaceType) {
            "toggle" -> {
                switch.visibility = View.VISIBLE
                switch.isChecked = robot.robotStatus.value

                switch.setOnCheckedChangeListener { _, isOn ->
                    onSwitch(isOn)
                }
            }
            "range" -> {
                seekBar.visibility = View.VISIBLE
                seekBar.max = robot.robotStatus.max
                seekBar.min = robot.robotStatus.min
                seekBar.progress = robot.robotStatus.current

                seekBar.setOnSeekBarChangeListener(object : SeekBar.OnSeekBarChangeListener {
                    override fun onProgressChanged(seekBar: SeekBar?, progress: Int, fromUser: Boolean) {
                        onSeek(progress)
                    }
                    override fun onStartTrackingTouch(seekBar: SeekBar?) {}
                    override fun onStopTrackingTouch(seekBar: SeekBar?) {}
                })
            }
        }
    }

    /**
     * onSwitch is called whenever the switch changes states
     *
     * @param isOn whether the switch is on/off
     */
    private fun onSwitch(isOn: Boolean) {
        Log.d(fTag, "onSwitch($isOn)")

        // Send the update to the server
        (activity as MainActivity).restApiService.robotToggle(robotId, isOn, mapOf()).enqueue(object: Callback<ResponseBody> {
            override fun onResponse(call: Call<ResponseBody>, response: Response<ResponseBody>) {
                if (response.isSuccessful) {
                    (activity as Context).toast("Success")
                    Log.d(fTag, "Server accepted setting switch to $isOn")
                } else {
                    val error = RestApiService.extractErrorFromResponse(response)
                    Log.e(fTag, "Getting the robot was unsuccessful, error: $error")
                    if (error != null) {
                        snackbar_layout.snackbar(error)
                    }
                }
            }

            override fun onFailure(call: Call<ResponseBody>, t: Throwable) {
                val error = t.message
                Log.e(fTag, "Robot get FAILED, error: $error")
                if (error != null) {
                    snackbar_layout.snackbar(error)
                }
            }
        })
    }

    /**
     * onSeek is called whenever the seek bar changes states
     *
     * @param value current of seek bar
     */
    private fun onSeek(value: Int) {
        Log.d(fTag, "onSeek($value)")
        // TODO: send update to server
    }

}