diff flyweight.go gofmt/flyweight.go
--- /var/folders/z8/wf90zpwx7fn290n5fs85xsd80000gn/T/gofmt253697928	2012-12-01 11:09:12.000000000 +0800
+++ /var/folders/z8/wf90zpwx7fn290n5fs85xsd80000gn/T/gofmt286532935	2012-12-01 11:09:12.000000000 +0800
@@ -3,8 +3,8 @@
 import "fmt"
 
 type Flyweight struct {
-    a string
-    b string
+    a   string
+    b   string
 }
 
 type FlyweightFactory struct {
@@ -12,14 +12,14 @@
 }
 
 func (p *FlyweightFactory) getFlyweight(a string, b string) *Flyweight {
-    key := a+b
+    key := a + b
     if _, present := p.pool[key]; !present {
-        p.pool[key] = &Flyweight{a,b}
+        p.pool[key] = &Flyweight{a, b}
     }
     return p.pool[key]
 }
 
-func compare (a, b *Flyweight) {
+func compare(a, b *Flyweight) {
     if a == b {
         fmt.Println(a, "and", b, "are the same object")
     } else {
