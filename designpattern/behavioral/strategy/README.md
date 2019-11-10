# Strategy

In Strategy pattern, a class behavior or its algorithm can be changed at run time. This type of design pattern comes under behavior pattern.

In Strategy pattern, we create objects which represent various strategies and a context object whose behavior varies as per its strategy object. The strategy object changes the executing algorithm of the context object.

# Actors

## Context

The Context maintains a reference to one of the concrete strategies and communicates with this object only via the strategy interface.

## Strategy

Strategy interface defining an action and concrete strategy classes implementing the Strategy interface.

## Concrete Strategy

Concrete Strategies implement different variations of an algorithm the context uses.

## Client

The Client creates a specific strategy object and passes it to the context. The context exposes a setter which lets clients replace the strategy associated with the context at runtime.

## Links

- https://refactoring.guru/design-patterns/strategy
- https://www.tutorialspoint.com/design_pattern/strategy_pattern.htm
- https://www.geeksforgeeks.org/strategy-pattern-set-1/
