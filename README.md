### Overview

Let's talk about how it works, simply this application is just a server that runs workers utilizing goroutine and the amount of worker is depending by how much user connected to server. This worker basically just standing by and read user's command then executes it.

### Command

```sh
telnet localhost 8888

```

for development, i'm using `telnet` with this command:

after you connected to the server, you can run these commands:

- `/nick [nickname]`
  sets your nickname, when freshly connected, you will be given a random name.

- `/join [room] [code]`
  joins to a room, when freshly connected, you'll be joining `lobby` room.

- `/msg [message]`
  broadcast a message to your room.

- `/rooms`
  list all available rooms.

- `/members`
  list all members in your room.

- `/whisper [target] [message]`
  whisper a message to one of the member in your room.

- `/private [room] [code]`
  creates a private room with code

### ToDo

1. Add graceful shutdown.
