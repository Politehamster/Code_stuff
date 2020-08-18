/*import java.util.concurrent.TimeUnit;

import org.openqa.selenium.WebDriver;
import org.openqa.selenium.chrome.ChromeDriver;

public class FirstTestInSelenium {

public static void main(String[] args) {
// TODO Auto-generated method stub

System.setProperty("webdriver.chrome.driver", ".\\Driver\\chromedriver.exe");
WebDriver driver=new ChromeDriver();
driver.manage().timeouts().implicitlyWait(10, TimeUnit.SECONDS);
driver.manage().window().maximize();
driver.get("https://www.prospector.at");
driver.close();

}

}*/

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/tebeka/selenium"
)


func main() {
	const (
		seleniumPath    = "h:/_Programs/_dev/Selenium_drivers/selenium-server-standalone-3.141.59.jar"
		geckoDriverPath = "h:/_Programs/_dev/Selenium_drivers/geckodriver.exe"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.GeckoDriver(geckoDriverPath),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
//		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
		panic(err)
	}

	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		panic(err)
	}
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	err = elem.SendKeys(`
		package main
		import "fmt"
		func main() {
			fmt.Println("Lolzies\n")
		}
	`)
	if err != nil {
		panic(err)
	}

	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

}
