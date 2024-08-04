## Tutorial

How to write "Hello world"

```scss
// main.scss

@use "fmt";

@func main() {
  $greeting = "Hello world";
  @import fmt.println($greeting);
}
```