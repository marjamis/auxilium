# Made with https://github.com/charmbracelet/vhs
Output demo.gif

Require go
Require auxilium

# Set up how the terminal looks
Set FontSize 30
Set Width 1500
Set Height 400

# Hiding the execution of the go file to start the demo
Hide
Type "auxilium --config ./test/demo.aux.yml"
Enter
Show

# Simulating the inputs of a person using the demo auxilium configuration. This is inline with the steps of the demo.aux.yml configuration
# Step 1
Sleep 3
Enter

# Step 2
Sleep 3
Enter

# Step 3
Sleep 3
Enter

# Step 4
Sleep 3
Enter

# Step 5
Type "echo In a shell..."
Enter
Sleep 2
Type "echo Exiting..."
Enter
Sleep 1
Type "exit"
Sleep 1
Enter
Enter

# Step 6
Sleep 3
Enter

# Step 7
Sleep 3
Enter

# Step 8
Sleep 3
Enter

# Step 9
Sleep 3
Enter

# Step 10
# Waiting a bit longer on the last step before stopping the gif
Sleep 5
