package io.github.kzmake.agwd.blog

import io.github.kzmake.agwd.blog.v1._
import scala.concurrent.{ExecutionContext, Future}

class Handler(implicit ec: ExecutionContext) extends BlogService {
  override def create(_request: CreateRequest): Future[CreateResponse] = {
    Future.successful(
      CreateResponse(
        Some(
          Blog(
            url = "http://example.com/bob"
          )
        )
      )
    )
  }
}
