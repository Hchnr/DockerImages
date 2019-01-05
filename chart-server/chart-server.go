package main
import (
    "bytes"
    "strconv"
    "os/exec"
    "net/http"
    "crypto/hmac"
    "crypto/sha1"
    "encoding/hex"
    "io/ioutil"

    "jiweil/chart-server/util"
    "jiweil/chart-server/common"
)

const (
    headerSignature = "X-Hub-Signature" // HTTP header where the sha1 signature of the payload is stored
)

var chr chan *http.Request

func exec_shell(s string) error {
    cmd := exec.Command("/bin/bash", "-c", s)

    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    return err
}

func do_update() {
    for req := range chr {
        util.Logger.Info("[Request path]:", req.URL.Path)
        util.Logger.Info("[Run shell]:", "./update.sh " + req.URL.Path)
        err := exec_shell("./update.sh " + req.URL.Path) 
        if err != nil {
            util.Logger.Error("Run shell error: ", err)
        } 
    }
}

func update(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    if r.Method != "POST" {
        w.Write([]byte("POST needed."))
        util.Logger.Info("[Request not POST]: ", r.Method)
        return
    }

    // read the HTTP request body
    payload, err := ioutil.ReadAll(r.Body)
    if err != nil {
        util.Logger.Info("[Read request error]:", err)
        return
    }

    // validate signature
    gitHubSecretToken := common.Config.Secret
    if gitHubSecretToken != "" {
        sign := r.Header.Get(headerSignature)

        // to compute the HMAC in order to check for equality with what has been sent by GitHub
        mac := hmac.New(sha1.New, []byte(gitHubSecretToken))
        mac.Write(payload)
        expectedHash := hex.EncodeToString(mac.Sum(nil))
        receivedHash := sign[5:] // remove 'sha1='

        // signature mismatch, do not process
        if !hmac.Equal([]byte(receivedHash), []byte(expectedHash)) {
            util.Logger.Info("[Request secret eror token]:", gitHubSecretToken)
            util.Logger.Info("[Request secret eror exp hash]:", expectedHash)
            util.Logger.Info("[Request secret eror rcv hash]:", receivedHash)
            return
        }
    }

    // ADD request to queue
    chr <- r
}

func health(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() 
    util.Logger.Info("[Request path]: update ", r.URL.Path ) 
}

func main() {
    chr = make(chan *http.Request, 100)

    http.Handle("/release/", http.StripPrefix("/release/", http.FileServer(http.Dir("./release"))))
    http.HandleFunc("/update", update)
    http.HandleFunc("/health", health)
    
    go do_update()
    http.ListenAndServe(":" + strconv.Itoa(common.Config.Port), nil)
}

