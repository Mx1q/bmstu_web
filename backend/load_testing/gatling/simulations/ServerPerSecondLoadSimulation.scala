import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class ServerPerSecondLoadSimulation extends Simulation {

  val httpProtocolNginx = http.baseUrl("http://nginx:80")
  val httpProtocolBack = http.baseUrl("http://backend-mirror:8081")

  val scnNginx = scenario("Nginx load test")
    .exec(http("Nginx Metrics").get("/api/v2/ingredients?page=1"))
  val scnBack = scenario("Backend load test")
      .exec(http("Backend Metrics").get("/api/v2/ingredients?page=1"))

  setUp(
    scnNginx.inject(
        constantUsersPerSec(100).during(30)
    ).protocols(httpProtocolNginx),
    scnBack.inject(
        constantUsersPerSec(100).during(30)
    ).protocols(httpProtocolBack)
  ).maxDuration(60.seconds)
}