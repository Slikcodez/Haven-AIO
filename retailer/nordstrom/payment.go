package nordstrom

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func generateNewURL(PAYPAGE_ID, Id, orderId, encryptedAccount, encryptedCvv, publicKey string) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := strconv.Itoa(int(math.Pow10(20) + rand.Float64()*(math.Pow10(21)-math.Pow10(20)-1)))
	timeStr := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	newURL := "https://request.eprotect.vantivcnp.com/eProtect/paypage?paypageId=" + PAYPAGE_ID +
		"&reportGroup=*merchant1500&id=" + Id +
		"&orderId=" + orderId +
		"&encryptedAccount=" + encryptedAccount +
		"&pciNonSensitive=false&encryptedCvv=" + encryptedCvv +
		"&publicKeyId=" + publicKey +
		"&requestType=eProtect&targetServer=primary&jsoncallback=JS" + randomNum + "_" + timeStr

	return newURL
}

func idStrings(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	randomStr := make([]rune, length)
	for i := range randomStr {
		randomStr[i] = letters[rand.Intn(len(letters))]
	}
	rand.Seed(time.Now().UnixNano())
	randomStr2 := make([]rune, length)
	for i := range randomStr2 {
		randomStr2[i] = letters[rand.Intn(len(letters))]
	}
	return (string(randomStr) + string(randomStr2))
}

func getModulusnKey() string {
	url := "https://your-url-here"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	body := string(bodyBytes)

	regex := regexp.MustCompile(`modulus:"(?P<modulus>[^"]+)",exponent:"(?P<exponent>[^"]+)",keyId:"(?P<keyId>[^"]+)",visaCheckoutApiKey:"(?P<apiKey>[^"]+)",visaCheckoutEncryptionKey:"(?P<encKey>[^"]+)"`)
	matches := regex.FindStringSubmatch(body)

	names := regex.SubexpNames()
	result := make(map[string]string)
	for i, match := range matches {
		result[names[i]] = match
	}

	/*fmt.Printf("Modulus: %s\n", result["modulus"])
	fmt.Printf("Exponent: %s\n", result["exponent"])
	fmt.Printf("KeyId: %s\n", result["keyId"])
	fmt.Printf("VisaCheckoutApiKey: %s\n", result["apiKey"])
	fmt.Printf("VisaCheckoutEncryptionKey: %s\n", result["encKey"])*/

	return fmt.Sprintf("%v:::%v", result["modulus"], result["keyId"])
}

func DoPayment() {
	PAYPAGE_ID := "dkZyzhtKZ2jpAa9T"
	Id := idStrings(8)
	orderId := idStrings(8)
	info := getModulusnKey()
	encryptedAccount := "your_encrypted_account"
	encryptedCvv := "your_encrypted_cvv"
	publicKey := "your_public_key"

	newURL := generateNewURL(PAYPAGE_ID, Id, orderId, encryptedAccount, encryptedCvv, publicKey)
	fmt.Println(newURL)
}
