
# g(o)root Exploitation Framework

General-purpose local exploitation framework for GNU/Linux. (WIP)


## Motivation

Most other languages suck when it comes to reusability of exploits.

The only thing that comes close to a reusable exploitation framework is
Metasploit, but Ruby sucks, too, and it's pretty much impossible to deploy
on target systems for any kind of killchain. I'm gonna spare you the
details of the proprietary corporate BS about ClownStrike.

The idea behind this framework is that it's intelligent enough for
automated privilege escalation purposes and that it iterates through
the most likely possible exploits incrementally, depending on the current
system's environment.


## Features (WIP)

- [x] Support for DirtyCow [CVE-2016-5195](https://nvd.nist.gov/vuln/detail/cve-2016-5195)
- [ ] Support for DirtyPipe [CVE-2022-0847](https://nvd.nist.gov/vuln/detail/cve-2022-0847)
- [ ] Support for DirtyEBPF [CVE-2022-23222](https://nvd.nist.gov/vuln/detail/cve-2022-23222)
- [ ] Support for Sudo Bypass [CVE-2019-14287](https://nvd.nist.gov/vuln/detail/CVE-2019-14287)
- [ ] Support for Mempod Dipper [CVE-2012-0056](https://nvd.nist.gov/vuln/detail/CVE-2012-0056)
- [ ] Support for Full Nelson [CVE-2010-4258](https://nvd.nist.gov/vuln/detail/CVE-2010-4258)
- [ ] Support for RDS [CVE-2010-3904](https://nvd.nist.gov/vuln/detail/CVE-2010-3904)
- [ ] Support for Sudo Inject
- [ ] Support for `LD_PRELOAD` injection of `SUID` binaries
- [ ] Support for `RPATH` injection of `SUID` binaries
- [ ] Support for [GTFOBins](https://gtfobins.github.io/)
- [ ] Support for writeable `/etc/cron.d` and `/etc/cron.{monthly,weekly,daily,hourly}`
- [ ] Support for writeable `/etc/NetworkManager/system-connections`
- [ ] Support for writeable `/etc/systemd/system`
- [ ] Support for writeable `/etc/passwd`
- [ ] Support for writeable `/etc/sudoers`


## TODO (WIP)

- [ ] Embedded x86 shellcode to spawn process
- [ ] Embedded x86 shellcode to spawn shell
- [ ] Embedded x64 shellcode to spawn process
- [ ] Embedded x64 shellcode to spawn shell


## License

AGPL3

(Bonus points if you use this in a web service)

