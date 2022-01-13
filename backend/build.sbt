import Settings._

val baseName = "agwd"

ThisBuild / organization               := "com.github.kzmake"
ThisBuild / scalaVersion               := "2.13.6"
ThisBuild / semanticdbEnabled          := true
ThisBuild / scalafixScalaBinaryVersion := CrossVersion.binaryScalaVersion(scalaVersion.value)
ThisBuild / semanticdbVersion          := scalafixSemanticdb.revision // only required for Scala 2.x
ThisBuild / scalacOptions ++= Seq(
  "-unchecked",
  "-deprecation",
  "-Wunused:imports" // Scala 2.x only, required by `RemoveUnused`
)

lazy val api = (project in file("api"))
  .settings(
    name := "api"
  )
  .enablePlugins(AkkaGrpcPlugin)
  .settings(protoSettings)
  .settings(
    Compile / PB.protoSources += file("../api")
  )

lazy val blog = (project in file("services/blog"))
  .settings(
    name                        := "blog",
    Docker / packageName        := "agwd/blog",
    Docker / dockerRepository   := Some("ghcr.io/kzmake"),
    Docker / maintainer         := "kzmake <kamake.i3a@gmail.com>",
    Docker / dockerExposedPorts := List(50051)
  )
  .enablePlugins(JavaAppPackaging)
  .settings(coreSettings, grpcSettings)
  .dependsOn(api)
  .aggregate(api)

lazy val user = (project in file("services/user"))
  .settings(
    name                        := "user",
    Docker / packageName        := "agwd/user",
    Docker / dockerRepository   := Some("ghcr.io/kzmake"),
    Docker / maintainer         := "kzmake <kamake.i3a@gmail.com>",
    Docker / dockerExposedPorts := List(50051)
  )
  .enablePlugins(JavaAppPackaging)
  .settings(coreSettings, grpcSettings)
  .dependsOn(api)
  .aggregate(api)

lazy val root = (project in file("."))
  .dependsOn(user, blog)
  .aggregate(user, blog)
