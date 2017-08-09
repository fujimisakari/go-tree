# Tree Command With Go Lang

Introduction
------------

I implemented tree commnad with go.  
We use yaml for tree definition.


Installation and usage
----------------------

```console
$ make setup
$ make build or make install
$ go-tree sample > ./sample.yaml

$ go-tree ./sample.yaml
Sample
├── fizz
├── buzz
├── dir1
│   ├── comp1-1
│   ├── comp1-2
│   ├── comp1-3
│   └── comp1-4
├── dir2
│   ├── comp2-1
│   ├── comp2-2
│   ├── comp2-3
│   ├── comp2-4
│   ├── dir3
│   │   ├── comp3-1
│   │   ├── comp3-2
│   │   ├── dir4
│   │   │   ├── comp4-1
│   │   │   ├── comp4-2
│   │   │   └── comp4-3
│   │   ├── comp3-3
│   │   └── comp3-4
│   ├── comp2-5
│   └── comp2-6
├── dir5
│   ├── comp4-1
│   ├── comp4-2
│   ├── comp4-3
│   └── comp4-4
├── foo
└── bar
```

About Tree Yaml
-------

For the definition of tree, from the `go-tree` onwards, define a tree structure with yaml's array. For directories, add a semicolon at the end. You can define the root directory name with `root-dir`(default: ".").

```yaml
root-dir: "Sample"
go-tree:
  - fizz
  - buzz
  - dir1:
    - comp1-1
    - comp1-2
    - comp1-3
    - comp1-4
  - dir2:
    - comp2-1
    - comp2-2
    - comp2-3
    - comp2-4
    - dir3:
      - comp3-1
      - comp3-2
      - dir4:
        - comp4-1
        - comp4-2
        - comp4-3
      - comp3-3
      - comp3-4
    - comp2-5
    - comp2-6
  - dir5:
    - comp4-1
    - comp4-2
    - comp4-3
    - comp4-4
  - foo
  - bar
```
