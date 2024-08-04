```scss
@use "fmt";

@mixin main() {
  $greeting = "Hello world";
  @import fmt.print-ln($greeting);
}
```
