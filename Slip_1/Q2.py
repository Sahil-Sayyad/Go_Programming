
# a. Draw block diagram /pin diagram of Raspberry-Pi/ Beagle board /Arduino Uno board interfacing with IR Sensor/Temperature Sensor/Camera.
#    (Internal Examiner assign any one option for board and interface device and respective interface programming option).

# b. WAP in python/C++ language to blink LED.
# c. Write down the observations on Input and Output.
# d. Write down the Result and Conclusion.



# b.  WAP in python language to blink LED.

import RPi.GPIO as GPIO # type: ignore # Import Respberry Pi GPIO library
import time             #Import time module
GPIO.setwarnings(False)
GPIO.setmode(GPIO.BCM)
#assign numbering for the GPIO using BCM
#GPIO.setmode(GPIO.BOARD)
#assign number for the GPIO using Board

cnt = 0
Blink_Time = 1  #change LED status every 1 seconds
RED_LED = 14
GPIO.setup(RED_LED, GPIO.OUT)
while True:
    if cnt == 0 :
        GPIO.output(RED_LED, False)
        cnt = 1
        
    else :
        GPIO.output(RED_LED, True)
        cnt = 0
        time.sleep(Blink_Time)
        GPIO.cleanup()
    

# c. Write down the observations on Input and Output.
 #Observations:
     #1. Input:
              #Any obstacle that comes in the range of the sensor is detected.IR sensors actually
              #measure the heat being emitted from the object. So the heat is the actual input for
              #sensors. The IR sensor gets its input from GPIO 21.
     #2. Output:
              #The output is shown by the LED's and the buzzer.When an obstacle is detected, the Green
              #LED glows and buzzer is turned on, i.e whenever heat is sensed by the sensor, output is shown.
              #The output is shown by making the GPIO 15(Green LED) and GPIO 18(Buzzer) HIGH.


# d. Write down the Result and Conclusion.
     #Result and Conclusion:
              #we have successfully implemented the connection of IR sensor with Raspberry Pi/Beagle bone/Arduino for obstacle detection.
              #The output is shown by a glowing LED's and buzzer.