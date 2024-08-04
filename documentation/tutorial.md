## Tutorial

### How to write "Hello world"

```scss

@use "fmt";

@func main() {
  $greeting = "Hello world";
  @import fmt.println($greeting);
}
```

### How to use for loop

```scss

@use "fmt";

@func main() {
  $sum = 0;

  @for $i from 1 through 3 {
    $sum += 1;
  }
  
  @import fmt.println($sum);
}
```
This will output: ```3```

### How to use while loop

```scss

@use "fmt";

@func main() {
  $sum = 0;

  @while $sum < 5 {
    $sum += 1;
  }
  
  @import fmt.println($sum);
}
```
This will output: ```5```
