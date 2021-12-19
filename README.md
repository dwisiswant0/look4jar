# Look4jar

![look4jar](https://user-images.githubusercontent.com/25837540/146642218-85adc53a-3df3-4ce8-b7a6-8268e4f2e03e.jpg)

Look for JAR files that vulnerable to [Log4j RCE](https://logging.apache.org/log4j/2.x/security.html) ([CVE‐2021‐44228](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-44228))

---

## Objectives

It differs from some other tools that scan for vulnerable remote services by running trigger exploits such as track DNS pingbacks. **Look4jar** tried to find `JndiLookup.class` file in the java archive _(recursively)_, if yet any — then it will look for `log4j2.enableJndiLookup` in `JndiManager.class` file which it deems possibly vulnerable.

## Installation

- Download a prebuilt binary from [releases page](https://github.com/dwisiswant0/look4jar/releases/latest), unpack and run! or:
- If you have **[Go1.16+](https://go.dev/doc/install)** compiler installed & configured:

```console
$ go install dw1.io/look4jar@latest
```

**— or**

Building from source code:

```console
$ git clone git@github.com:dwisiswant0/look4jar.git
$ cd look4jar/
$ go mod tidy
$ go build .
$ ./look4jar -h
```

## Usage

```

  look4jar
  ---
  Look for JAR files that vulnerable to Log4j RCE (CVE‐2021‐44228)
  @dwisiswant0


Usage:
  look4jar -p /path/to/file [OPTIONS...]

Options:
  -p, --path <FILE/PATH>      Specify EAR/JAR/WAR file/directory to scan recursively
  -v, --verbose               Verbose mode (default false)

Examples:
  look4jar -p /path/to/file.jar
  look4jar -p /usr/local/lib -v
```

## Similar Projects

- [yahoo/check-log4j](https://github.com/yahoo/check-log4j)

## License

**Look4jar** is distributed under Apache License v2.0. See `LICENSE`.