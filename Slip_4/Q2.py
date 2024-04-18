'''
a. Draw block diagram /pin diagram of Raspberry-Pi/ Beagle board/Arduino Uno board interfacing with IR Sensor/TemperatureSensor/Camera.
    (Internal Examiner assign any one option for board and interface device and respective interface programming option).

b. WAP in python/C++ language to toggle two LED’s.
c. Write down the observations on Input and Output.
d. Write down the Result and Conclusion.

'''

# b. WAP in python/C++ language to toggle two LED’s.

import RPi.GPIO as GPIO  #import Raspberry Pi GPIO library from time import sleep
import time as sleep     #Import the sleep function from the time module.
GPIO.setwarnings(False)  #Ignore warning for now 
GPIO.setmode(GPIO.BOARD) #Use physical pin numbering 
RED_LED = 14 
GREEN_LED = 15
GPIO.setup(RED_LED, GPIO.OUT, initial=GPIO.LOW)
GPIO.setup(GREEN_LED, GPIO.OUT, initial=GPIO.LOW)

while True :              #Run forever
    GPIO.output(RED_LED, True) # Turn ON
    GPIO.output(GREEN_LED, False) # Turn OFF
    sleep(1)               #Sleep for 1 second
    GPIO.output(RED_LED,False) #Turn OFF 
    GPIO.output(GREEN_LED, True) #Turn On
    sleep(1)  #Sleep for 1 second


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