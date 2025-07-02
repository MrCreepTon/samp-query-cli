# samp-query-cli
😍 CLI to query SAMP servers directly from terminal written in golang!

👉 Usage: `samp-query [IP:Port] <-r to show rules>`

Example output for `samp-query cheat.hackmysoftware.ru -r`:
<pre>
>> Basic Info <<
Hostname: HackMySoftware - Cheating Server | Version: 4.1
Address: cheat.hackmysoftware.ru:7777
Ping: 925 ms
Players: 3 / 50
Gamemode: Unknown/HMS/Cheating/HackMySoftware
>> Rules <<
worldtime: 12:00
allowed_clients: 0.3.7, 0.3.DL
artwork: No
lagcomp: On
mapname: Cheating Town
version: omp 1.2.0.2670
weather: 10
weburl: vk.com/hackmysoftware
</pre>

⚠️ If no `-r` set, rules section wont be visible.
