package main
import(
    "testing"
)
func TestMain(t *testing.T){
    got:=Hello()
    want:="breh"
    if got!=want{
        t.Errorf("got %q want %q",got,want)
    }
}
