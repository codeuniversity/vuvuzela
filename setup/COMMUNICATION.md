# Communicating through a mesh network

## Sending messages from a slave to a master

**Note: Although this is a nice way to see that the mesh network is actually working, we're not going to utilize 
a master-slave system for our robots. For our use-case, it's better to have a mesh network without hierarchies.**

### Starting the server

To send messages between nodes of the network, the easiest way is to use Alfred, a communication server for mesh
networks (especially for batman-adv, which we are using). 

Assuming that you have two Raspberry Pis ready and used the `run_me.sh` script to set up the network, you should have
Alfred already installed on both of them. You now need to start a master server on **one** of the Pis by using the following
command:

`sudo alfred -i wlan0 -b bat0 -m`

 On the other Pi, start a slave server by using the same command without the `-m`:
 
 `sudo alfred -i wlan0 -b bat0`
 
 ### Sending messages
 
 Open a new terminal instance on both of the Pis so you don't have to stop the server on either of them. 
 
 To send a message as a slave, you need to assign an ID to the message (to make them easily distinguishable). IDs 0-63
 are reserved by Alfred, so we'll pick 64:
 
 `echo "This is a test message" | sudo alfred -s 64`
 
Switch to the other Pi (the master node) and run the following command to retrieve all messages with ID 64:

`sudo alfred -r 64`

Your output should look like this:

```
{ "b8:27:eb:64:b7:94", "This is a test message\x0a" },
```

Congrats, you successfully transmitted a message within a mesh network!

## Sending messages without hierarchy

For now, it is sufficient to spin up a master server for each and every node. This way, they will synchronize all data
among all nodes. I will investigate further whether this technique has any major downsides.

## More information

To really dive deep into this topic, I found the documentation of the
[Alfred architecture](https://www.open-mesh.org/projects/alfred/wiki/Alfred_architecture) quite helpful.