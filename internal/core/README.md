# `/core`

Contains the domain model and business logic, which are technology-agnostic.

1. `models`: Represent domain-specific business objects and encapsulate business rules.
2. `app`: Contain application-specific business rules and orchestrate the flow of data between entities.
3. `ports`: Define the interaction points between the core and the external world.
   These are abstract interfaces that allow the core to communicate with the outside world.