# isro-launching
There are a couple of constraints that ISRO needs to work under:

1. The group of 4 satellites that need to be launched from each Launch Site needs to happen simultaneously with each satellite counting down(10...0) and launching. 

2. The launch of the Satellite group happens in tandem. What it means is, SATs 1-4 launched from LS1 needs to notify LS2 of the successful launch and SATs 5-8 then are launched from LS2. The same drill happens, LS2 notifies LS1 of the launch and LS1 would launch SATs 9-12. This happens until till the time all the satellites are launched. Please build a GO program to get make the launch and the communication happen. Unit Tests for the code would be great. Please upload the program on GitHub and send me the instructions to check out the code and run it on our machines.

Time Complexity: O(n)
