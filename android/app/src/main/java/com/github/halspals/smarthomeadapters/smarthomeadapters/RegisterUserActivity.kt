package com.github.halspals.smarthomeadapters.smarthomeadapters

import android.content.Context
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.view.View
import android.view.inputmethod.InputMethodManager
import com.github.halspals.smarthomeadapters.smarthomeadapters.model.Token
import com.github.halspals.smarthomeadapters.smarthomeadapters.model.User
import kotlinx.android.synthetic.main.activity_register_user.*
import org.jetbrains.anko.*
import org.jetbrains.anko.design.snackbar
import org.json.JSONException
import org.json.JSONObject
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class RegisterUserActivity : AppCompatActivity() {

    private val tag = "RegisterUserActivity"

    private val restApiService by lazy {
        RestApiService.new()
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_register_user)

        /*
            Set up the click events for the buttons.
         */
        register_button.setOnClickListener { _ ->

            // Dismiss the keyboard
            val inputMethodManager = getSystemService(Context.INPUT_METHOD_SERVICE) as InputMethodManager
            val currentView = currentFocus ?: View(this)
            inputMethodManager.hideSoftInputFromWindow(currentView.windowToken, 0)

            // Signal to the user that we are waiting for a remote call
            register_button.isEnabled = false
            progressBar.visibility = View.VISIBLE

            val email = email_input.text.toString()
            val password = password_input.text.toString()
            val confirmedPassword = confirm_password_input.text.toString()

            val inputsOK = checkEmailInput(email)
                    && checkPasswordInput(password)
                    && checkConfirmPasswordInput(password, confirmedPassword)

            if (inputsOK) {
                registerNewUser(User(email, password))
            } else {
                register_button.isEnabled = true
                progressBar.visibility = View.GONE
            }
        }


        login_button.setOnClickListener { _ ->
            // Return to the AuthenticationActivity which launched this RegisterUserActivity
            finish()
        }

        /*
            Set onFocusChangeListeners for the input fields to check if their input thus far
            is valid.
         */
        email_input.setOnFocusChangeListener { _, hasFocus ->
            if (!hasFocus) {
                checkEmailInput(email_input.text.toString())
            }
        }

        password_input.setOnFocusChangeListener { _, hasFocus ->
            if (!hasFocus) {
                checkPasswordInput(password_input.text.toString())
            }
        }

        confirm_password_input.setOnFocusChangeListener { _, hasFocus ->
            if (!hasFocus) {
                checkConfirmPasswordInput(
                        password_input.text.toString(), confirm_password_input.text.toString())
            }
        }


    }

    /**
     * Makes a REST call to register the given user.
     *
     * @param user the user to register an account for
     *
     */
    private fun registerNewUser(user: User) {
        restApiService.registerUser(user).enqueue(object : Callback<User> {
            override fun onResponse(call: Call<User>, response: Response<User>) {
                val responseUser = response.body()
                if (response.isSuccessful) {
                    assert(user == responseUser) {
                        "Returned user doesn't equal expected user. " +
                                "Expected {$user}, received {$responseUser}"
                    }
                    toast("Registered; now signing you in...")
                    signInUser(user)
                } else {
                    val error = RestApiService.extractErrorFromResponse(response)
                    handleCallbackError(error, enableFurtherRegistration = true)
                }
            }

            override fun onFailure(call: Call<User>, error: Throwable) {
                handleCallbackError(error.message, enableFurtherRegistration = true)
            }
        })
    }

    /**
     * Makes a REST call to sign in the given user.
     *
     * @param user the user to sign in
     */
    private fun signInUser(user: User) {
        restApiService.loginUser(user).enqueue(object: Callback<Token> {
            override fun onResponse(call: Call<Token>, response: Response<Token>) {
                val token = response.body()
                if (response.isSuccessful && token != null) {
                    saveTokenAndMoveToMain(token.token)
                } else {
                    val error = RestApiService.extractErrorFromResponse(response)
                    handleCallbackError(error)
                }
            }

            override fun onFailure(call: Call<Token>, error: Throwable) {
                handleCallbackError(error.message)
            }
        })
    }

    /**
     * WIP: Saves the token and starts a [MainActivity].
     * Currently only saves the token in shared prefs, unencrypted.
     *
     * @param token the authorization token received from the server
     */
    private fun saveTokenAndMoveToMain(token: String) {
        Log.d(tag, "Succeeded in receiving token, saving and starting MainActivity")

        val prefs = getSharedPreferences(SHARED_PREFERENCES_FILE, Context.MODE_PRIVATE)
        val editor = prefs.edit()
        editor.putString(ACCESS_TOKEN_KEY, token)
        editor.apply()

        toast("Signed in")
        startActivity(intentFor<MainActivity>(ACCESS_TOKEN_KEY to token).clearTask().newTask())
    }

    /**
     * Handles an error received by an api call, displaying the message to the user.
     *
     * @param error the error received from the api call
     * @param whether to allow the user to press register again
     */
    private fun handleCallbackError(error: String?, enableFurtherRegistration: Boolean = false) {

        // Display the error to the user
        Log.d(tag, "Login failed: $error")
        if (error != null) {
            snackbar_layout.snackbar(error)
        } else {
            Log.w(tag, "Error message was null")
        }

        register_button.isEnabled = enableFurtherRegistration
        progressBar.visibility = View.GONE
    }

    /**
     * Makes sure the given email is valid, setting an error to [email_input] if not.
     *
     * @param email the email to check for validity
     * @return whether the given email is a valid one
     */
    private fun checkEmailInput(email: String): Boolean {

        val emailIsValid = android.util.Patterns.EMAIL_ADDRESS.matcher(email).matches()

        if (!emailIsValid) {
            email_input.error = "Invalid email address"
        } else {
            email_input.error = null
        }

        return emailIsValid
    }


    /**
     * Checks the validity of the password (>= 8 chars).
     * Sets an error to [password_input] if it is not valid.
     *
     * @param password the password to check
     * @return whether the password is a valid one
     */
    private fun checkPasswordInput(password: String): Boolean {
        // TODO we can add more checks here like password length etc
        val pwIsValid = password.length >= 8

        if (!pwIsValid) {
            password_input.error = "Must be at least 8 characters long"
        } else {
            password_input.error = null
        }

        return pwIsValid
    }

    /**
     * Checks that the two passwords match, adding an error to [confirm_password_input] if not.
     *
     * @return whether the two passwords match
     */
    private fun checkConfirmPasswordInput(password: String, confirmed_password: String): Boolean {
        // Simply make sure that the two passwords given match.

        val confirmedPwIsValid = password == confirmed_password

        if (!confirmedPwIsValid) {
            confirm_password_input.error = "Passwords do not match"
        } else {
            confirm_password_input.error = null
        }

        return confirmedPwIsValid
    }
}
