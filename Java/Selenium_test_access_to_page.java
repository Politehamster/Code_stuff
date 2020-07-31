import java.util.concurrent.TimeUnit;

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

}