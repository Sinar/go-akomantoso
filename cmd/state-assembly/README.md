# State Assembly

## Build

```
make cli
```

## Usage

Assumption: Raw files is ./raw/StateAssembly/Hansard/ directory

PlanIt
```
$  ./cmd/state-assembly/go-akn planit --source HANSARD-15-JULAI-2020.pdf

```

Extract into SayIt format
```
$  ./cmd/state-assembly/go-akn sayit --source HANSARD-15-JULAI-2020.pdf
```