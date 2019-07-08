// package fsrepo
//
// TODO explain the package roadmap...
//
//   .ipws/
//   ├── client/
//   |   ├── client.lock          <------ protects client/ + signals its own pid
//   │   ├── ipws-client.cpuprof
//   │   └── ipws-client.memprof
//   ├── config
//   ├── daemon/
//   │   ├── daemon.lock          <------ protects daemon/ + signals its own address
//   │   ├── ipws-daemon.cpuprof
//   │   └── ipws-daemon.memprof
//   ├── datastore/
//   ├── repo.lock                <------ protects datastore/ and config
//   └── version
package fsrepo

// TODO prevent multiple daemons from running
