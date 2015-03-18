# gcd

Recurse through $HOME/src/go/src, finds the first file / directory that matches the given argument and prints it.

e.g.:
```
$ gcd gcd
/Users/raphael/src/go/src/github.com/raphael/gcd
```

## Usage
In your `.bashrc`, `.bash_profile` or equivalent:
```bash
function cdg() {
   cd `gcd $1`
}
```
Then:
```
~ $ cdg gdc
~ src/go/src/github.com/raphael/gcd
```
