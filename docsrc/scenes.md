# Scene transition

Here's scene transitions.

```uml
@startuml

(*)      --> "title"
"title"  --> "game"
"game"   --> "result"
"result" --> "title"

@enduml
```

## title

* First scene of this application. Show title.
* Tap anywhere to go next scene.
* Show highest score from past game plays.

* TODO: put picture

## game

* Main game scene.
* Basically endless.
* Tap "Finish" to go next scene.

* TODO: put picture

## result

* Show score of latest game play.
* Save score if it is high enough.
* Tap "Back to title" to go back to "title"

* TODO: put picture

