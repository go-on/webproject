/*
 TODO: make watchers for

 - less compilation
 - typescript compilation
 - compilation in go project,
    if a file is changed and is a go file,
        - the package of the dir of the file is compiled
        - the tests are run
        - and the main package is compiled (all dirs up until a main.go is found or GOPATH is reached)          
- for css classes and ids: compile them into the typescript and less directory
  they should be included in the other files
*/

package main

import (
	"gopkg.in/fsnotify.v1"
	"fmt"
	"gopkg.in/metakeule/goh4.v5"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

//	"sync/atomic"
)

var (
//cmdId = atomic.A
)

// sync.WaitGroup

type (
	Compiler interface {
		Compile(path string) (string, error)
		Name() string
		Lock()
		Unlock()
		// directories to watch
		Dirs() []string
		Affected(path string) bool
	}

	Notifyer interface {
		Notify(msg string)
	}

	compiler struct {
		*sync.Mutex
		Bin       string
		Options   []string
		Dir       string
		Ext       string
		name      string
		Ignore    []string // ignored directories as regular expressions
		watchDirs []string
	}

	less struct {
		*compiler
		OutPutFile string
		MainFile   string
	}

	typescript struct {
		*compiler
		OutPutFile string
	}

	gocompiler struct {
		*compiler
		GoPath   string
		MainFile string
	}

	ProjectWatcher struct {
		*sync.Mutex
		Compilers []Compiler
		Watcher   *fsnotify.Watcher
		Notifyer  Notifyer
		Pool      map[Compiler]string
		Frequency time.Duration
		Ready     chan int
	}
)

type execStruct struct {
	output string
	err    error
}

var Debug = false

func Exec(name string, opts ...string) (output string, err error) {
	//ch := make(chan execStruct, 1)
	//go func() {
	stamp := time.Now().UnixNano()
	id := fmt.Sprintf("%s %s [%v]", name, strings.Join(opts, " "), stamp)
	if Debug {
		log.Printf("executing: %s", id)
	}
	cmd := exec.Command(name, opts...)
	var out []byte
	out, err = cmd.CombinedOutput()
	output = string(out)
	if Debug {
		log.Printf("output of %s:\n%s", id, output)
	}
	if err != nil {
		err = fmt.Errorf(output)
		log.Printf("error of %s:\n%s", id, err.Error())
	}
	// ch <- execStruct{outp, e}
	//}()
	//x := <-ch
	//output = x.output
	//err = x.err
	return
}

func Which(cmd string) (path string, err error) {
	path, err = Exec("which", cmd)
	path = strings.TrimRight(path, "\n")
	if path == "" {
		err = fmt.Errorf("not found: %s", cmd)
	}
	return
}

func NewGoCompiler(mainfile string, dir string, ignore []string, options ...string) Compiler {
	if mainfile != "" {
		if filepath.Ext(mainfile) != ".go" {
			log.Fatalf("mainfile %#v is not a go file", mainfile)
		}
	}
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("GOPATH is not set")
	}

	bin, err := Which("go")
	if err != nil {
		log.Fatal("you do not have the go compiler installed.")
	}
	/*
		if len(options) == 0 {
			// recommended options 
			//options := []string{"--strict-imports", "--verbose", "--no-color", "--line-numbers=all", "--strict-math=on", "--strict-units=off"}
		}
	*/
	return &gocompiler{NewCompiler(dir, "go build", ".go", bin, ignore, options...), gopath, mainfile}
}

func (ø *gocompiler) Compile(file string) (output string, err error) {
	if file != ø.MainFile {
		dir := filepath.Dir(file)
		packageName := strings.Replace(dir, path.Join(ø.GoPath, "src")+"/", "", 1)
		output, err = Exec(ø.compiler.Bin, "install", packageName)
		if err != nil {
			return
		}
	}

	if ø.MainFile == "" {
		return
	}

	// mainDir := filepath.Dir(ø.MainFile)
	//packageNameMain := strings.Replace(mainDir, path.Join(ø.GoPath, "src")+"/", "", 1)
	//mainOutput := strings.Replace(ø.MainFile, ".go", "", 1)
	//_, err = Exec(ø.compiler.Bin, "build", "-o="+mainOutput, packageNameMain)
	//output, err = Exec(ø.compiler.Bin, "build", "-o="+mainOutput, ø.MainFile)
	//if err != nil {
	//output, err = Exec("/bin/sh", "-c='"+mainOutput+" & '")
	//eChan := make (chan error, 1)
	//output, err = Exec(ø.compiler.Bin, "run", ø.MainFile)
	go Exec(ø.compiler.Bin, "run", ø.MainFile)
	//}
	return
}

func NewCompiler(dir string, name string, ext string, bin string, ignore []string, options ...string) (ø *compiler) {
	ø = &compiler{
		Mutex:     &sync.Mutex{},
		Bin:       bin,
		Options:   options,
		Dir:       dir,
		Ext:       ext,
		name:      name,
		Ignore:    ignore,
		watchDirs: []string{},
	}
	var err error
	ø.watchDirs, err = AllAffectedDirs(ø.Dir, ø.Ignore)
	if err != nil {
		log.Fatalf("can't set watching dirs: %s", err.Error())
	}
	ø.watchDirs = append(ø.watchDirs, ø.Dir)
	return
}

func IsIgnored(file string, ignores []string) (is bool, err error) {
	for _, i := range ignores {
		reg, e := regexp.Compile(i)
		if e != nil {
			err = e
			return
		}
		if reg.MatchString(file) {
			is = true
			return
		}
	}
	return
}

func AllAffectedDirs(base string, ignores []string) (all []string, err error) {
	// TODO for every subdir take the affected dirs and merge them
	dirs, err := ioutil.ReadDir(base)
	if err != nil {
		return
	}
	all = []string{}
	for _, dir := range dirs {
		if dir.IsDir() == false {
			continue
		}
		is, e := IsIgnored(dir.Name(), ignores)
		if e != nil {
			err = e
			return
		}
		if is {
			continue
		}
		dirPath := path.Join(base, dir.Name())
		all = append(all, dirPath)
		var subs []string
		subs, err = AllAffectedDirs(dirPath, ignores)
		if err != nil {
			return
		}
		all = append(all, subs...)
	}
	return
}

func (ø *compiler) Dirs() []string { return ø.watchDirs }

func (ø *compiler) Name() string { return ø.name }

func (ø *compiler) Compile(file string) (output string, err error) {
	opts := append([]string{}, ø.Options...)
	opts = append(opts, file)
	output, err = Exec(ø.Bin, opts...)
	return
}

func (ø *compiler) Affected(file string) bool {
	if filepath.Ext(file) == ø.Ext && IsInDir(file, ø.Dir) {
		return true
	}
	return false
}

func NewLess(mainFile string, dir string, outputFile string, ignore []string, options ...string) Compiler {
	bin, err := Which("lessc")
	if err != nil {
		log.Fatal("you do not have the less compiler installed. please run 'npm install -g less'")
	}
	if len(options) == 0 {
		// recommended options 
		options = []string{"--strict-imports", "--verbose", "--no-color", "--line-numbers=all", "--strict-math=on", "--strict-units=off"}
	}
	return &less{NewCompiler(dir, "less (lessc)", ".less", bin, ignore, options...), outputFile, mainFile}
}

func CompileToFile(file string, output string) error {
	return ioutil.WriteFile(file, []byte(output), os.FileMode(0644))
}

func (ø *less) Compile(file string) (output string, err error) {
	opts := append([]string{}, ø.compiler.Options...)
	opts = append(opts, ø.MainFile)
	output, err = Exec(ø.Bin, opts...)
	if err != nil {
		return
	}
	err = CompileToFile(ø.OutPutFile, output)
	return
}

/*
func (ø *typescript) Compile(file) (output string, err error) {
	output, err = ø.compiler.Compile(file)
	if err != nil {
		return
	}
	err = CompileToFile(ø.OutPutFile, output)
	return
}
*/

func NewTypeScript(dir string, outputDir string, ignore []string, options ...string) Compiler {
	bin, err := Which("tsc")
	if err != nil {
		log.Fatal("you do not have the typescript compiler installed. please run 'npm install -g typescript'")
	}
	if len(options) == 0 {
		// recommended options 
		options = []string{"--out", outputDir, "-c", "--disallowbool", "--disallowbool", "--sourcemap", "--module", "amd"}
	}
	return &typescript{NewCompiler(dir, "typescript (tsc)", ".ts", bin, ignore, options...), outputDir}
}

func IsInDir(path string, dir string) bool {
	return strings.Contains(path, dir)
}

func New(notifyer Notifyer, compilers ...Compiler) (ø *ProjectWatcher) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("can't create watcher: %s\n", err.Error())
	}
	return &ProjectWatcher{
		Mutex:     &sync.Mutex{},
		Compilers: compilers,
		Watcher:   watcher,
		Ready:     make(chan int, 1),
		Notifyer:  notifyer,
		Frequency: time.Millisecond * 20,
		// each compiler should be handled at one time with the last file succeeding
		Pool: map[Compiler]string{},
	}
}

func (ø *ProjectWatcher) SendMessages() {
	for {
		ø.Lock()
		for comp, file := range ø.Pool {
			log.Printf("handle %s with %s\n", file, comp.Name())
			out, err := comp.Compile(file)
			if err != nil {
				ø.Notifyer.Notify(out)
			}
			delete(ø.Pool, comp)
		}
		ø.Unlock()
		time.Sleep(ø.Frequency)
	}
}

func (ø *ProjectWatcher) HandleFile(path string) {
	// log.Printf("trying to handle: %s\n", path)
	for _, c := range ø.Compilers {
		c.Lock()
		if c.Affected(path) {
			ø.Pool[c] = path

		}
		c.Unlock()
	}
}

func (ø *ProjectWatcher) Run() (err error) {
	for _, c := range ø.Compilers {
		for _, d := range c.Dirs() {
			err = ø.Watcher.Watch(d)
			if err != nil {
				return
			}
		}
	}

	go ø.SendMessages()

	go func() {
		for {
			select {
			case ev := <-ø.Watcher.Event:
				//log.Println("event:", ev)

				what := ""
				handleIt := true
				switch {
				case ev.IsCreate():
					what = "created"
					handleIt = false
				case ev.IsDelete():
					handleIt = false
					what = "deleted"
				case ev.IsModify():
					what = "modified"
				case ev.IsRename():
					handleIt = false
					what = "renamed"
				}
				_ = what
				//log.Println("file: ", ev.Name, " ", what)

				if handleIt {
					ø.Lock()
					ø.HandleFile(ev.Name)
					ø.Unlock()
				}

			case err := <-ø.Watcher.Error:
				log.Println("watcher error:", err)
			}
		}
		ø.Lock()
		ø.Ready <- 1
		ø.Unlock()
	}()
	return
}

type zenityBin string

func (ø zenityBin) Notify(msg string) {
	Exec(string(ø), "--info", fmt.Sprintf("--text=%s", msg))
}

func Zenity() zenityBin {
	path, err := Which("zenity")
	if err != nil {
		log.Fatalf("can't find zenity")
	}
	return zenityBin(path)
}

type stripGoPath struct {
	Notifyer
	GoPath string
}

func StripGoPath(n Notifyer) Notifyer {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("GOPATH is not set")
	}
	return &stripGoPath{n, path.Join(gopath, "src") + "/"}
}

func (ø *stripGoPath) Notify(msg string) {
	nm := strings.Replace(msg, ø.GoPath, "", -1)
	ø.Notifyer.Notify(nm)
}

/*
	        Flags
		        FSN_CREATE = 1
		        FSN_MODIFY = 2
		        FSN_DELETE = 4
		        FSN_RENAME = 8

		        FSN_ALL = FSN_MODIFY | FSN_DELETE | FSN_RENAME | FSN_CREATE
*/
//watcher.WatchFlags(path string, flags uint32) error {

// watcher.RemoveWatch(path string) 

//watcher.Close()

type cssVarsCompiler struct {
	*sync.Mutex
	All        []goh4.Class
	Package    string
	Dir        string
	OutputFile string
	Bin        string
	name       string
	Var        string
}

func (ø *cssVarsCompiler) Name() string              { return "cssVars compiler (" + ø.name + ")" }
func (ø *cssVarsCompiler) Dirs() []string            { return []string{ø.Dir} }
func (ø *cssVarsCompiler) Affected(path string) bool { return IsInDir(path, ø.Dir) }
func (ø *cssVarsCompiler) Compile(path string) (output string, err error) {
	// -in=koelnart/frontend/style/class -target=/home/benny/Entwicklung/gopath/src/koelnart/frontend/less/classes.less
	if ø.Var == "" {
		output, err = Exec(ø.Bin, "-in="+ø.Package, "-target="+ø.OutputFile)
	} else {
		output, err = Exec(ø.Bin, "-in="+ø.Package, "-target="+ø.OutputFile, "-var="+ø.Var)
	}
	return
}

func NewCssVars(name string, dir string, output string, var_ string) (ø *cssVarsCompiler) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("GOPATH is not set")
	}

	var err error
	ø = &cssVarsCompiler{}
	ø.name = name
	ø.Bin, err = Which("cssvars")
	if err != nil {
		log.Fatal("you do not have cssvars installed. run go get github.com/metakeule/goh4/css/cssvars")
	}

	/*
		if !IsInDir(gopath, dir) {
			log.Fatalf("%s is not in $GOPATH", dir)
		}
	*/
	d, e := os.Stat(dir)
	if e != nil {
		log.Fatalf("%s does not exists: %s", dir, e.Error())
	}
	if !d.IsDir() {
		log.Fatalf("%s is no package", dir)
	}
	ø.Dir = dir
	packageName := strings.Replace(dir, path.Join(gopath, "src")+"/", "", 1)
	ø.Package = packageName
	ø.Var = var_
	ø.OutputFile = output
	ø.Mutex = &sync.Mutex{}
	return
}

/*
	Compiler interface {
		Compile(path string) (string, error)
		Name() string
		Lock()
		Unlock()
		// directories to watch
		Dirs() []string
		Affected(path string) bool
	}
*/

func CompilersForApp(baseDir string, app string) []Compiler {
	appDir := path.Join(baseDir, "app", app)
	staticDir := path.Join(baseDir, "static")
	lessDir := path.Join(appDir, "less")
	lessOutput := path.Join(staticDir, "css", app, "all.css")
	lessMain := path.Join(lessDir, "main.less")
	typeScriptDir := path.Join(appDir, "typescript")
	typeScriptOutput := path.Join(staticDir, "js", app)
	classDir := path.Join(appDir, "class")
	idDir := path.Join(appDir, "id")
	classOutput := path.Join(lessDir, "classes.less")
	idOutput := path.Join(lessDir, "ids.less")

	return []Compiler{
		NewCssVars(app+" classes.less", classDir, classOutput, ""),
		NewCssVars(app+" ids.less", idDir, idOutput, ""),
		NewLess(lessMain, lessDir, lessOutput, []string{}),
		NewTypeScript(typeScriptDir, typeScriptOutput, []string{}),
	}

}

func CompilersForAllApps(baseDir string) (cs []Compiler) {
	dirs, e := ioutil.ReadDir(path.Join(baseDir, "app"))
	if e != nil {
		log.Fatalf("could not read dir %s: %s", baseDir, e.Error())
	}
	cs = []Compiler{}
	for _, dir := range dirs {
		if dir.IsDir() == true {
			cs = append(cs, CompilersForApp(baseDir, dir.Name())...)
		}
	}
	return
}

func CompilersForProject(baseDir string) (compilers []Compiler) {
	mainFile := path.Join(baseDir, "main.go")
	compilers = []Compiler{
		NewGoCompiler(mainFile, baseDir, []string{"static", "less", "typescript"}),
	}
	compilers = append(compilers, CompilersForAllApps(baseDir)...)
	return
}

func main() {
	compilers := CompilersForProject("/home/benny/Entwicklung/gopath/src/koelnart")
	projectWatcher := New(StripGoPath(Zenity()), compilers...)
	projectWatcher.Run()
	<-projectWatcher.Ready
	for {
	}
}
