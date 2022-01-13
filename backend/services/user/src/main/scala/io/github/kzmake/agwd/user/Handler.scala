package io.github.kzmake.agwd.user

import io.github.kzmake.agwd._
import akka.actor.typed.ActorSystem
import akka.grpc.internal.{GrpcProtocolNative, Identity}
import akka.grpc.GrpcClientSettings
import scala.concurrent.duration._
import scala.concurrent.ExecutionContext
import scala.concurrent.Future
import scala.util.Success
import scala.util.Failure

class Handler(system: ActorSystem[_])(implicit ec: ExecutionContext) extends user.v1.UserService {
  implicit val sys: ActorSystem[_] = system
  implicit val writer              = GrpcProtocolNative.newWriter(Identity)

  override def create(
      _request: user.v1.CreateRequest
  ): Future[user.v1.CreateResponse] = {
    val client = blog.v1
      .BlogServiceClient(
        GrpcClientSettings
          .connectToServiceAt("localhost", 50001)
          .withDeadline(10.second)
          .withTls(false)
      )
    val req = new blog.v1.CreateRequest()
    val reply = client
      .create()
      .addHeader("dapr-app-id", "blog")
      .invoke(req)

    reply.transform {
      case Success(msg) =>
        Success(
          user.v1.CreateResponse(
            Some(
              user.v1.User(
                id = "12345",
                name = "alice",
                blogUrl = msg.getBlog.url,
                emailAddress = "alice@example.com",
                walletId = "1234"
              )
            )
          )
        )
      case Failure(e) => Failure(e)
    }
  }
}
