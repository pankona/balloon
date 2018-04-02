# Modules and sequences

Here's Modules' responsibilities and relations.

## Basic event loop

```uml
@startuml

hide footbox

participant "Main loop\n(Drive)" as main
participant "Event\nQueue"       as eq
participant "Event\nSubscribers" as es

loop while game scene
  main -> eq : fetch events
  main <- eq : events
  loop len(events)
    main -> es : fire events
    es -> es : do something
    note right of es : each events may\npush new event to\nevent queue
    eq <- es : (may push a new events)
  end
end

@enduml
```

* It is necessary to care the amount of events not to slow down game loop.
  * All events fired within a Drive function call,  
  should be completed within 1/60 sec if the game want to keep 60 fps.

## Strategy of balloon spawning

* Basically ballon spawns in random. Following factors are in random.
  * Spawn interval.
  * Balloon appears from bottom line but X position.
  * Balloon's moving speed.
  * Balloon's size.
  * Balloon's color.
  * A frog is in balloon or not.
* But obey following rules.
  * Don't spawn multiple balloons in same time.

```uml
@startuml

hide footbox

participant "Main loop\n(Drive)" as m
participant "Event\nQueue"       as eq
participant Spawner              as s
participant Factory              as f

== initialize ==
create s
m -> s : new Spawner(frame counter)
m <- s : event subscribe
create f
m -> f : new Factory
m <- f : event subscribe
====
m -> s : run
eq <- s : push(JudgeSpawn)
loop while game scene
  m -> eq : fetch events
  m <- eq : events
  loop len(events)
  == event is JudgeSpawn ==
    m -> s : event publish
    m -> f : event publish
    note right of f : ignore
    alt
      note right of s : push event\nin random
      s -> eq : push(Spawn)
    end
    s -> eq : push(JudgeSpawn)
  == event is Spawn ==
    m -> s : event publish
    note right of s : ignore
    m -> f : event publish
    f -> f : spawn()
    note right of f : spawn a new balloon
  end
end

@enduml
```

## modules

* modules are connected by event queue.

```uml
class eventQueue {
}

class game {
}

class spawner {
}

class factory {
}

class balloon {
}

class scorer {
}

eventQueue --> game    : pub
eventQueue --> spawner : pub
eventQueue --> factory : pub
eventQueue --> balloon : pub
eventQueue --> scorer  : pub
eventQueue --> control : pub

game    --> eventQueue : sub
spawner --> eventQueue : sub
factory --> eventQueue : sub
balloon --> eventQueue : sub
scorer  --> eventQueue : sub
control --> eventQueue : sub
```

### modules description

* eventQueue
  * Use to publish/subscribe events by game module
  * Consists of a single channel

* game
  * Handle user input (tap or drag)
  * Handle user interaction for buttons

* spawner
  * Judge whether spawn a new balloon or not
  * loops in some time interval

* factory
  * Spawns a new balloon by receiving Spawn event
  * Has responsibility to determine where to spawn

* balloon
  * Spawned by factory.
  * Balloons will be bursted by tapping or dragging
  * Bursting balloon will send a Bust event
  * Balloons will send Escape event by disappearing from screen
  * Run away of balloon will send a Escape event

* scorer
  * Increase/decrease user's score by receiving Burst/Escape event

## events

| event      | publisher | subscriber           | next event (possible) | description                                                                                              |
|------------|-----------|----------------------|-----------------------|----------------------------------------------------------------------------------------------------------|
| JudgeSpawn | spawner   | spawner              | JudgeSpawn<br>Spawn   | Judge whether spawn a new balloon (frog) or not                                                          |
| Spawn      | spawner   | factory              | Spawned               | A new balloon will be spawned                                                                            |
| Spawned    | factory   | -                    | -                     | A new balloon has been spawned                                                                           |
| Shoot      | game      | balloon              | -                     | Shoot action by tapping screen                                                                           |
| Slash      | game      | balloon              | -                     | Slash action by dragging screen                                                                          |
| Burst      | balloon   | scorer<br>JudgeSpawn | -                     | A balloon (or frog) is bursted by shoot or slash.<br>scorer increases score according to what is bursted |
| Escape     | balloon   | scorer               | -                     | A balloon (or frog) is run away from screen                                                              |
| Finish     | control   | game                 | -                     | Finish game by tapping Finish button                                                                     |


