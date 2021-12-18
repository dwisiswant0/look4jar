package main

const (
  POSSIBLE_VULN = "Possibly vulnerable '%s'"
  JNDI_ENABLED  = "log4j2.enableJndiLookup"
  JAR_PATTERN   = "\\.[ejw]ar$"
  FOUND_JAR     = "Found JAR file: %s"

  ERR_OPEN = "Failed to open '%s' file"
  ERR_LOOK = "Can't look up for '%s' path: %s"

  MSG_BANNER = `
  look4jar
  ---
  Look for JAR files that vulnerable to Log4j RCE (CVE‐2021‐44228)
  @dwisiswant0

`
  MSG_USAGE = `
Usage:
  look4jar -p /path/to/file [OPTIONS...]

Options:
  -p, --path <FILE/PATH>      Specify EAR/JAR/WAR file/directory to scan recursively
  -v, --verbose               Verbose mode (default false)

Examples:
  look4jar -p /path/to/file.jar
  look4jar -p /usr/local/lib -v
`
)
