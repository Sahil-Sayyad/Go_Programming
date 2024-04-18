'''
a. Draw block diagram /pin diagram of Raspberry-Pi/ Beagle board/Arduino Uno board interfacing with IR Sensor/TemperatureSensor/Camera.
     (Internal Examiner assign any one option for board and interface device and respective interface programming option).

b. WAP in python/C++ language to turn ON/OFF buzzer.
c. Write down the observations on Input and Output.
d. Write down the Result and Conclusion.

'''

# b. WAP in python/C++ language to turn ON/OFF buzzer.

import RPi.GPIO as GPIO # import Raspberry Pi GPIO library from time import sleep.
import time as sleep
#import the sleep function from the time module
GPIO.setwarnings(False) #ignore wanrning for now
GPIO.setmode(GPIO.BOARD) # Use physical pin numbering
Buzzer = 18
GPIO.setup(Buzzer,GPIO.OUT, initial=GPIO.LOW)

while True:         #Run forever
    GPIO.output(Buzzer, True) # turn ON
    sleep(1) # Sleep for 1 Second
    GPIO.output(Buzzer,False) # turn OFF
    sleep(1) #Sleep for 1 Second



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