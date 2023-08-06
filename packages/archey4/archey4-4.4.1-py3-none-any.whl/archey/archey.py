#!/usr/bin/env python3


import json
import os
import re
import sys
from enum import Enum
from glob import glob
from os import getenv, getuid
from subprocess import CalledProcessError, DEVNULL, PIPE, Popen, \
    TimeoutExpired, check_output


# ----- Distributions fingerprints ---- #

class Distributions(Enum):
    ARCH_LINUX = 'Arch.*'
    BUNSENLABS = 'BunsenLabs'
    CRUNCHBANG = 'CrunchBang'
    DEBIAN = '(Rasp|De)bian'
    FEDORA = 'Fedora'
    GENTOO = 'Gentoo'
    KALI_LINUX = 'Kali'
    MANJARO_LINUX = 'ManjaroLinux'
    LINUX = 'Linux'
    LINUX_MINT = 'LinuxMint'
    OPENSUSE = 'openSUSE'
    RED_HAT = 'Red Hat'
    UBUNTU = 'Ubuntu'
    WINDOWS = 'Windows'


# ------------ Dictionaries ----------- #

colorDict = {
    Distributions.ARCH_LINUX: ['\x1b[0;34m', '\x1b[1;34m'],
    Distributions.BUNSENLABS: ['\x1b[1;37m', '\x1b[1;33m', '\x1b[0;33m'],
    Distributions.CRUNCHBANG: ['\x1b[1;37m', '\x1b[1;37m'],
    Distributions.DEBIAN: ['\x1b[0;31m', '\x1b[1;31m'],
    Distributions.FEDORA: ['\x1b[1;37m', '\x1b[0;34m'],
    Distributions.GENTOO: ['\x1b[1;37m', '\x1b[1;35m'],
    Distributions.KALI_LINUX: ['\x1b[1;37m', '\x1b[1;34m'],
    Distributions.MANJARO_LINUX: ['\x1b[0;32m', '\x1b[1;32m'],
    Distributions.LINUX: ['\x1b[1;33m', '\x1b[1;37m'],
    Distributions.LINUX_MINT: ['\x1b[1;37m', '\x1b[1;32m'],
    Distributions.OPENSUSE: ['\x1b[1;37m', '\x1b[1;32m'],
    Distributions.RED_HAT: ['\x1b[1;37m', '\x1b[1;31m', '\x1b[0;31m'],
    Distributions.UBUNTU: ['\x1b[0;31m', '\x1b[1;31m', '\x1b[0;33m'],
    Distributions.WINDOWS: ['\x1b[1;31m', '\x1b[1;34m',
                            '\x1b[1;32m', '\x1b[0;33m'],
    'sensors': ['\x1b[0;32m', '\x1b[0;33m', '\x1b[0;31m'],
    'clear': '\x1b[0m'
}

deDict = {
    'cinnamon': 'Cinnamon',
    'gnome-session': 'GNOME',
    'gnome-shell': 'GNOME',
    'mate-session': 'MATE',
    'ksmserver': 'KDE',
    'xfce4-session': 'Xfce',
    'fur-box-session': 'Fur Box',
    'lxsession': 'LXDE',
    'lxqt-session': 'LXQt'
}

wmDict = {
    'awesome': 'Awesome',
    'beryl': 'Beryl',
    'blackbox': 'Blackbox',
    'bspwm': 'bspwm',
    'cinnamon': 'Cinnamon',
    'compiz': 'Compiz',
    'dwm': 'DWM',
    'enlightenment': 'Enlightenment',
    'herbstluftwm': 'herbstluftwm',
    'fluxbox': 'Fluxbox',
    'fvwm': 'FVWM',
    'i3': 'i3',
    'icewm': 'IceWM',
    'kwin_x11': 'KWin',
    'metacity': 'Metacity',
    'musca': 'Musca',
    'openbox': 'Openbox',
    'pekwm': 'PekWM',
    'qtile': 'QTile',
    'ratpoison': 'ratpoison',
    'scrotwm': 'ScrotWM',
    'stumpwm': 'StumpWM',
    'subtle': 'Subtle',
    'monsterwm': 'MonsterWM',
    'wingo': 'Wingo',
    'wmaker': 'Window Maker',
    'wmfs': 'Wmfs',
    'wmii': 'wmii',
    'xfwm4': 'Xfwm',
    'xmonad': 'xmonad'
}

logosDict = {
    Distributions.ARCH_LINUX: """
{c[1]}               +                 {r[0]}
{c[1]}               #                 {r[1]}
{c[1]}              ###                {r[2]}
{c[1]}             #####               {r[3]}
{c[1]}             ######              {r[4]}
{c[1]}            ; #####;             {r[5]}
{c[1]}           +##.#####             {r[6]}
{c[1]}          +##########            {r[7]}
{c[1]}         ######{c[0]}#####{c[1]}##;          {r[8]}
{c[1]}        ###{c[0]}############{c[1]}+         {r[9]}
{c[1]}       #{c[0]}######   #######{c[1]}         {r[10]}
{c[1]}     {c[0]}.######;     ;###;`\".{c[1]}       {r[11]}
{c[1]}    {c[0]}.#######;     ;#####.{c[1]}        {r[12]}
{c[1]}    {c[0]}#########.   .########`{c[1]}      {r[13]}
{c[1]}   {c[0]}######'           '######{c[1]}     {r[14]}
{c[1]}  {c[0]};####                 ####;{c[1]}    {r[15]}
{c[1]}  {c[0]}##'                     '##{c[1]}    {r[16]}
{c[1]} {c[0]}#'                         `#{c[1]}   {r[17]}\n""",
    Distributions.BUNSENLABS: """
{c[0]}           .{c[1]}..{c[0]}+hhy+-`         \0
{c[0]}        `+hd{c[1]}-{c[0]}+dddd{c[2]}hyso{c[0]}+//:    \0
{c[0]}      `+dddd{c[1]}:-{c[0]}sdh/.           {r[0]}
{c[0]}     -hdddddh{c[1]}-.{c[2]}/:{c[0]}             {r[1]}
{c[0]}    /ddddddddd{c[1]}:```{c[0]}            {r[2]}
{c[0]}   :ddddddddddd/              {r[3]}
{c[0]}  `hdddddddddddd+             {r[4]}
{c[0]}  /dddddddddddddd:            {r[5]}
{c[0]}  odddds..sddddddh            {r[6]}
{c[0]}  oddd/    /dddddd:           {r[7]}
{c[0]}  +dd+      +ddddd+           {r[8]}
{c[0]}  .dd`      `ddddd+           {r[9]}
{c[0]}   oh        ydddd:           {r[10]}
{c[0]}   `o        sdddh`           {r[11]}
{c[0]}             yddd:            {r[12]}
{c[0]}            `dddo             {r[13]}
{c[0]}     :s     :dds              {r[14]}
{c[0]}     yd/    yd+               {r[15]}
{c[0]}   `sddy   :h-                {r[16]}
{c[0]}  `sddys`  :`                 {r[17]}
{c[0]} -hdy+`y+yo./+/:-             \0
{c[0]} ...  .o++oso+/               \n""",
    Distributions.CRUNCHBANG: """
{c[0]}                 ___       ___      _    {r[0]}
{c[0]}                /  /      /  /     | |   {r[1]}
{c[0]}               /  /      /  /      | |   {r[2]}
{c[0]}              /  /      /  /       | |   {r[3]}
{c[0]}      _______/  /______/  /______  | |   {r[4]}
{c[0]}     /______   _______   _______/  | |   {r[5]}
{c[0]}           /  /      /  /          | |   {r[6]}
{c[0]}          /  /      /  /           | |   {r[7]}
{c[0]}         /  /      /  /            | |   {r[8]}
{c[0]}  ______/  /______/  /______       | |   {r[9]}
{c[0]} /_____   _______   _______/       | |   {r[10]}
{c[0]}      /  /      /  /               | |   {r[11]}
{c[0]}     /  /      /  /                |_|   {r[12]}
{c[0]}    /  /      /  /                  _    {r[13]}
{c[0]}   /  /      /  /                  | |   {r[14]}
{c[0]}  /__/      /__/                   |_|   {r[15]}
{c[0]}                                         {r[16]}
{c[0]}                                         {r[17]}\n""",
    Distributions.DEBIAN: """
{c[1]}                                  {r[0]}
{c[1]}          _sudZUZ#Z#XZo=_         {r[1]}
{c[1]}       _jmZZ2!!~---~!!X##wx       {r[2]}
{c[1]}    .<wdP~~            -!YZL,     {r[3]}
{c[1]}   .mX2'       _xaaa__     XZ[.   {r[4]}
{c[1]}   oZ[      _jdXY!~?S#wa   ]Xb;   {r[5]}
{c[1]}  _#e'     .]X2(     ~Xw|  )XXc   {r[6]}
{c[1]} .2Z`      ]X[.       xY|  ]oZ(   {r[7]}
{c[1]} .2#;      )3k;     _s!~   jXf`   {r[8]}
{c[1]}  {c[0]}1Z>{c[1]}      {c[0]}-]Xb/{c[1]}    {c[0]}~{c[1]}    {c[0]}__#2({c[1]}    {r[9]}
{c[1]}  {c[0]}-Zo;{c[1]}       {c[0]}+!4ZwerfgnZZXY'{c[1]}      {r[10]}
{c[1]}   {c[0]}*#[,{c[1]}        {c[0]}~-?!!!!!!-~{c[1]}        {r[11]}
{c[1]}    {c[0]}XUb;.{c[1]}                         {r[12]}
{c[1]}     {c[0]})YXL,,{c[1]}                       {r[13]}
{c[1]}       {c[0]}+3#bc,{c[1]}                     {r[14]}
{c[1]}         {c[0]}-)SSL,,{c[1]}                  {r[15]}
{c[1]}            {c[0]}~~~~~{c[1]}                 {r[16]}
{c[1]}                                  {r[17]}\n""",
    Distributions.FEDORA: """
{c[1]}           :/------------://          {r[0]}
{c[1]}        :------------------://        {r[1]}
{c[1]}      :-----------{c[0]}/shhdhyo/{c[1]}-://       {r[2]}
{c[1]}    /-----------{c[0]}omMMMNNNMMMd/{c[1]}-:/      {r[3]}
{c[1]}   :-----------{c[0]}sMMMdo:/{c[1]}       -:/     {r[4]}
{c[1]}  :-----------{c[0]}:MMMd{c[1]}-------    --:/    {r[5]}
{c[1]}  /-----------{c[0]}:MMMy{c[1]}-------    ---/    {r[6]}
{c[1]} :------    --{c[0]}/+MMMh/{c[1]}--        ---:   {r[7]}
{c[1]} :---     {c[0]}oNMMMMMMMMMNho{c[1]}     -----:   {r[8]}
{c[1]} :--      {c[0]}+shhhMMMmhhy++{c[1]}   ------:    {r[9]}
{c[1]} :-      -----{c[0]}:MMMy{c[1]}--------------/    {r[10]}
{c[1]} :-     ------{c[0]}/MMMy{c[1]}-------------:     {r[11]}
{c[1]} :-      ----{c[0]}/hMMM+{c[1]}------------:      {r[12]}
{c[1]} :--{c[0]}:dMMNdhhdNMMNo{c[1]}-----------:        {r[13]}
{c[1]} :---{c[0]}:sdNMMMMNds:{c[1]}----------:          {r[14]}
{c[1]} :------{c[0]}:://:{c[1]}-----------://           {r[15]}
{c[1]} :--------------------://             {r[16]}
{c[1]}                                      {r[17]}\n""",
    Distributions.GENTOO: """
{c[1]}          -/oyddmdhs+:.                {r[0]}
{c[1]}      -o{c[0]}dNMMMMMMMMNNmhy+{c[1]}-`             {r[1]}
{c[1]}    -y{c[0]}NMMMMMMMMMMMNNNmmdhy{c[1]}+-           {r[2]}
{c[1]}  `o{c[0]}mMMMMMMMMMMMMNmdmmmmddhhy{c[1]}/`        {r[3]}
{c[1]}  om{c[0]}MMMMMMMMMMMN{c[1]}hhyyyo{c[0]}hmdddhhhd{c[1]}o`      {r[4]}
{c[1]} .y{c[0]}dMMMMMMMMMMd{c[1]}hs++so/s{c[0]}mdddhhhhdm{c[1]}+`    {r[5]}
{c[1]}  oy{c[0]}hdmNMMMMMMMN{c[1]}dyooy{c[0]}dmddddhhhhyhN{c[1]}d.   {r[6]}
{c[1]}   :o{c[0]}yhhdNNMMMMMMMNNNmmdddhhhhhyym{c[1]}Mh   {r[7]}
{c[1]}     .:{c[0]}+sydNMMMMMNNNmmmdddhhhhhhmM{c[1]}my   {r[8]}
{c[1]}        /m{c[0]}MMMMMMNNNmmmdddhhhhhmMNh{c[1]}s:   {r[9]}
{c[1]}     `o{c[0]}NMMMMMMMNNNmmmddddhhdmMNhs{c[1]}+`    {r[10]}
{c[1]}   `s{c[0]}NMMMMMMMMNNNmmmdddddmNMmhs{c[1]}/.      {r[11]}
{c[1]}  /N{c[0]}MMMMMMMMNNNNmmmdddmNMNdso{c[1]}:`        {r[12]}
{c[1]} +M{c[0]}MMMMMMNNNNNmmmmdmNMNdso{c[1]}/-           {r[13]}
{c[1]} yM{c[0]}MNNNNNNNmmmmmNNMmhs+/{c[1]}-`             {r[14]}
{c[1]} /h{c[0]}MMNNNNNNNNMNdhs++/{c[1]}-`                {r[15]}
{c[1]} `/{c[0]}ohdmmddhys+++/:{c[1]}.`                   {r[16]}
{c[1]} `-//////:--.                          {r[17]}\n""",
    Distributions.KALI_LINUX: """
{c[0]}      ,.....                                        \0
{c[0]}  ----`     `..,;:ccc,.                             \0
{c[0]}           ......''';lxO.                           {r[0]}
{c[0]} .....''''..........,:ld;                           {r[1]}
{c[0]}            .';;;:::;,,.x,                          {r[2]}
{c[0]}       ..'''.            0Xxoc:,.  ...              {r[3]}
{c[0]}   ....                ,ONkc;,;cokOdc',.            {r[4]}
{c[0]}  .                   OMo           ':{c[1]}d{c[0]}o.           {r[5]}
{c[0]}                     dMc               :OO;         {r[6]}
{c[0]}                     0M.                 .:o.       {r[7]}
{c[0]}                     ;Wd                            {r[8]}
{c[0]}                      ;XO,                          {r[9]}
{c[0]}                        ,d0Odlc;,..                 {r[10]}
{c[0]}                            ..',;:cdOOd::,.         {r[11]}
{c[0]}                                     .:d;.':;.      {r[12]}
{c[0]}                                        'd,  .'     {r[13]}
{c[0]}                                          ;l   ..   {r[14]}
{c[0]}                                           .o       {r[15]}
{c[0]}                                             c      {r[16]}
{c[0]}                                             .'     {r[17]}
{c[0]}                                              .     \n""",
    Distributions.MANJARO_LINUX: """
{c[0]} $$$$$$$$$$$$$$$$  $$$$$$$   {r[0]}
{c[0]} M77777777777777M  M77777M   {r[1]}
{c[0]} M77777777777777M  M77777M   {r[2]}
{c[0]} M77777MMMMMMMMMM  M77777M   {r[3]}
{c[0]} M77777M           M77777M   {r[4]}
{c[0]} M77777M  $$$$$$$  M77777M   {r[5]}
{c[0]} MMMMMMM  M77777M  M77777M   {r[6]}
{c[0]}          M77777M  M77777M   {r[7]}
{c[0]} $$$$$$$  M77777M  M77777M   {r[8]}
{c[0]} M77777M  M77777M  M77777M   {r[9]}
{c[0]} M77777M  M77777M  M77777M   {r[10]}
{c[0]} M77777M  M77777M  M77777M   {r[11]}
{c[0]} M77777M  M77777M  M77777M   {r[12]}
{c[0]} M77777M  M77777M  M77777M   {r[13]}
{c[0]} M77777M  M77777M  M77777M   {r[14]}
{c[0]} M77777M  M77777M  M77777M   {r[15]}
{c[0]} M77777M  M77777M  M77777M   {r[16]}
{c[0]} MMMMMMM  MMMMMMM  MMMMMMM   {r[17]}\n""",
    Distributions.LINUX: """
{c[1]}                           {r[0]}
{c[1]}          a8888b.          {r[1]}
{c[1]}         d888888b.         {r[2]}
{c[1]}         8P"YP"Y88         {r[3]}
{c[1]}         8|o||o|88         {r[4]}
{c[1]}         8{c[0]}\\vvvv/{c[1]}88         {r[5]}
{c[1]}         8{c[0]} \\vv/ {c[1]}Y8.        {r[6]}
{c[1]}        d/  {c[0]}`'{c[1]}  \8b.       {r[7]}
{c[1]}      .dP   .     Y8b.     {r[8]}
{c[1]}     d8:'   "   `::88b.    {r[9]}
{c[1]}    d8"           `Y88b    {r[10]}
{c[1]}   :8P     '       :888    {r[11]}
{c[1]}    8a.    :      _a88P    {r[12]}
{c[1]}  {c[0]}._/"{c[1]}Yaa_ :    .{c[0]}| {c[1]}88P{c[0]}|{c[1]}    {r[13]}
{c[1]} {c[0]}\++++{c[1]}YP"      `{c[0]}| {c[1]}8P{c[0]}++\.{c[1]}   {r[14]}
{c[1]} {c[0]}/+++++\.{c[1]}_____.d{c[0]}|+++++/{c[1]}    {r[15]}
{c[1]}  {c[0]}\++++++){c[1]}888888P{c[0]}\+++/{c[1]}     {r[16]}
{c[1]}                           {r[17]}\n""",
    Distributions.LINUX_MINT: """
{c[1]}                                       {r[0]}
{c[1]} MMMMMMMMMMMMMMMMMMMMMMMMMmds+.        {r[1]}
{c[1]} MMm----::-://////////////oymNMd+`     {r[2]}
{c[1]} MMd      {c[0]}/++{c[1]}                -sNMd:    {r[3]}
{c[1]} MMNso/`  {c[0]}dMM{c[1]}    {c[0]}`.::-. .-::.`{c[1]} .hMN:   {r[4]}
{c[1]} ddddMMh  {c[0]}dMM{c[1]}   {c[0]}:hNMNMNhNMNMNh:`{c[1]} NMm   {r[5]}
{c[1]}     NMm  {c[0]}dMM{c[1]}  {c[0]}.NMN/-+MMM+-/NMN`{c[1]} dMM   {r[6]}
{c[1]}     NMm  {c[0]}dMM{c[1]}  {c[0]}-MMm{c[1]}  {c[0]}`MMM{c[1]}   {c[0]}dMM.{c[1]} dMM   {r[7]}
{c[1]}     NMm  {c[0]}dMM{c[1]}  {c[0]}-MMm{c[1]}  {c[0]}`MMM{c[1]}   {c[0]}dMM.{c[1]} dMM   {r[8]}
{c[1]}     NMm  {c[0]}dMM{c[1]}  {c[0]}.mmd{c[1]}  {c[0]}`mmm{c[1]}   {c[0]}yMM.{c[1]} dMM   {r[9]}
{c[1]}     NMm  {c[0]}dMM`{c[1]}  {c[0]}..`{c[1]}  {c[0]}`...{c[1]}   {c[0]}ydm.{c[1]} dMM   {r[10]}
{c[1]}     hMM-  {c[0]}+MMd/-------...-:sdds{c[1]} MMM   {r[11]}
{c[1]}     -NMm-  {c[0]}:hNMNNNmdddddddddy/`{c[1]} dMM   {r[12]}
{c[1]}      -dMNs-``{c[0]}-::::-------.``{c[1]}    dMM   {r[13]}
{c[1]}       `/dMNmy+/:-------------:/yMMM   {r[14]}
{c[1]}          ./ydNMMMMMMMMMMMMMMMMMMMMM   {r[15]}
{c[1]}                                       {r[16]}
{c[1]}                                       {r[17]}\n""",
    Distributions.OPENSUSE: """
{c[0]}            .;ldkO0000Okdl;.              {r[0]}
{c[0]}        .;d00xl:^''''''^:ok00d;.          {r[1]}
{c[0]}      .d00l'                'o00d.        {r[2]}
{c[0]}    .d0Kd'  {c[1]}Okxol:;,.{c[0]}          :O0d.      {r[3]}
{c[0]}   .OK{c[1]}KKK0kOKKKKKKKKKKOxo:,{c[0]}      lKO.     {r[4]}
{c[0]}  ,0K{c[1]}KKKKKKKKKKKKKKK0P^{c[0]},,,{c[1]}^dx:{c[0]}    ;00,    {r[5]}
{c[0]} .OK{c[1]}KKKKKKKKKKKKKKKk'{c[0]}.oOPPb.{c[1]}'0k.{c[0]}   cKO.   {r[6]}
{c[0]} :KK{c[1]}KKKKKKKKKKKKKKK: {c[0]}kKx..dd{c[1]} lKd{c[0]}   'OK:   {r[7]}
{c[0]} dKK{c[1]}KKKKKKKKKOx0KKKd {c[0]}^0KKKO'{c[1]} kKKc{c[0]}   dKd   {r[8]}
{c[0]} dKK{c[1]}KKKKKKKKKK;.;oOKx,..{c[0]}^{c[1]}..;kKKK0.{c[0]}  dKd   {r[9]}
{c[0]} :KK{c[1]}KKKKKKKKKK0o;...^cdxxOK0O/^^'{c[0]}  .0K:   {r[10]}
{c[0]}  kKK{c[1]}KKKKKKKKKKKKK0x;,,......,;od{c[0]}  lKk    {r[11]}
{c[0]}  '0K{c[1]}KKKKKKKKKKKKKKKKKKKK00KKOo^{c[0]}  c00'    {r[12]}
{c[0]}   'kK{c[1]}KKOxddxkOO00000Okxoc;''{c[0]}   .dKk'     {r[13]}
{c[0]}     l0Ko.                    .c00l'      {r[14]}
{c[0]}      'l0Kk:.              .;xK0l'        {r[15]}
{c[0]}         'lkK0xl:;,,,,;:ldO0kl'           {r[16]}
{c[0]}             '^:ldxkkkkxdl:^'             {r[17]}\n""",
    Distributions.RED_HAT: """
{c[1]}                                             {r[0]}
{c[1]}              {c[2]}\`.-..........\`{c[1]}               {r[1]}
{c[1]}             {c[2]}\`////////::.\`-/.{c[1]}              {r[2]}
{c[1]}             {c[2]}-: ....-////////.{c[1]}               {r[3]}
{c[1]}             {c[2]}//:-::///////////\`{c[1]}             {r[4]}
{c[1]}      {c[2]}\`--::: \`-://////////////:{c[1]}            {r[5]}
{c[1]}      {c[2]}//////-    \`\`.-://///////{c[1]} .\`        {r[6]}
{c[1]}      {c[2]}\`://////:-.\`    :///////::///:\`{c[1]}     {r[7]}
{c[1]}        {c[2]}.-/////////:---/////////////:{c[1]}        {r[8]}
{c[1]}           {c[2]}.-://////////////////////.{c[1]}        {r[9]}
{c[1]}          {c[0]}yMN+\`.-${c[2]}::///////////////-\`{c[1]}      {r[10]}
{c[1]}       {c[0]}.-\`:NMMNMs\`  \`..-------..\`{c[1]}        {r[11]}
{c[1]}        {c[0]}MN+/mMMMMMhoooyysshsss{c[1]}               {r[12]}
{c[1]} {c[0]}MMM    MMMMMMMMMMMMMMyyddMMM+{c[1]}               {r[13]}
{c[1]}  {c[0]}MMMM   MMMMMMMMMMMMMNdyNMMh\`     hyhMMM{c[1]}   {r[14]}
{c[1]}   {c[0]}MMMMMMMMMMMMMMMMyoNNNMMM+.   MMMMMMMM{c[1]}     {r[15]}
{c[1]}    {c[0]}MMNMMMNNMMMMMNM+ mhsMNyyyyMNMMMMsMM{c[1]}      {r[16]}
{c[1]}                                             {r[17]}\n""",
    Distributions.UBUNTU: """
{c[1]}                           {c[0]}.oyhhs:{c[1]}     {r[0]}
{c[1]}                  ..--.., {c[0]}shhhhhh-{c[1]}     {r[1]}
{c[1]}                -+++++++++`:{c[0]}yyhhyo`{c[1]}    {r[2]}
{c[1]}           {c[2]}.--{c[1]}  -++++++++/-.-{c[0]}::-`{c[1]}      {r[3]}
{c[1]}         {c[2]}.::::-{c[1]}   :-----:/+++/++/.     {r[4]}
{c[1]}        {c[2]}-:::::-.{c[1]}          .:++++++:    {r[5]}
{c[1]}   ,,, {c[2]}.:::::-`{c[1]}             .++++++-   {r[6]}
{c[1]} ./+++/-{c[2]}`-::-{c[1]}                ./////:   {r[7]}
{c[1]} +++++++ {c[2]}.::-{c[1]}                          {r[8]}
{c[1]} ./+++/-`{c[2]}-::-{c[1]}                {c[0]}:yyyyyo{c[1]}   {r[9]}
{c[1]}   ``` `{c[2]}-::::-`{c[1]}             {c[0]}:yhhhhh:{c[1]}   {r[10]}
{c[1]}       {c[2]}-:::::-.{c[1]}          {c[0]}`-ohhhhhh+{c[1]}    {r[11]}
{c[1]}         {c[2]}.::::-`{c[1]} {c[0]}-o+///+oyhhyyyhy:{c[1]}     {r[12]}
{c[1]}          {c[2]}`.--{c[1]}  {c[0]}/yhhhhhhhy+{c[2]},....{c[1]}       {r[13]}
{c[1]}                {c[0]}/hhhhhhhhh{c[2]}-.-:::;{c[1]}      {r[14]}
{c[1]}                {c[0]}`.:://::- {c[2]}-:::::;{c[1]}      {r[15]}
{c[1]}                          {c[2]}`.-:-'{c[1]}       {r[16]}
{c[1]}                                       {r[17]}\n""",
    Distributions.WINDOWS: """
{c[1]}                                        {r[0]}
{c[1]}         {c[0]},.=:^!^!t3Z3z.,{c[1]}                {r[1]}
{c[1]}        {c[0]}:tt:::tt333EE3{c[1]}                  {r[2]}
{c[1]}        {c[0]}Et:::ztt33EEE{c[1]}  {c[2]}@Ee.,{c[1]}      {c[2]}..,{c[1]}   {r[3]}
{c[1]}       {c[0]};tt:::tt333EE7{c[1]} {c[2]};EEEEEEttttt33#{c[1]}   {r[4]}
{c[1]}      {c[0]}:Et:::zt333EEQ.{c[1]} {c[2]}SEEEEEttttt33QL{c[1]}   {r[5]}
{c[1]}      {c[0]}it::::tt333EEF{c[1]} {c[2]}@EEEEEEttttt33F{c[1]}    {r[6]}
{c[1]}     {c[0]};3=*^```'*4EEV{c[1]} {c[2]}:EEEEEEttttt33@.{c[1]}    {r[7]}
{c[1]}     ,.=::::it=.,{c[0]} `{c[1]} {c[2]}@EEEEEEtttz33QF{c[1]}     {r[8]}
{c[1]}    ;::::::::zt33){c[1]}   {c[2]}'4EEEtttji3P*{c[1]}      {r[9]}
{c[1]}   :t::::::::tt33.{c[3]}:Z3z..{c[2]}  `` {c[3]},..g.{c[1]}      {r[10]}
{c[1]}   i::::::::zt33F{c[1]} {c[3]}AEEEtttt::::ztF{c[1]}       {r[11]}
{c[1]}  ;:::::::::t33V{c[1]} {c[3]};EEEttttt::::t3{c[1]}        {r[12]}
{c[1]}  E::::::::zt33L{c[1]} {c[3]}@EEEtttt::::z3F{c[1]}        {r[13]}
{c[1]} {{3=*^```'*4E3){c[1]} {c[3]};EEEtttt:::::tZ`{c[1]}        {r[14]}
{c[1]}             `{c[1]} {c[3]}:EEEEtttt::::z7{c[1]}          {r[15]}
{c[1]}                 {c[3]}'VEzjt:;;z>*`{c[1]}          {r[16]}
{c[1]}                                        {r[17]}\n"""
}


# ----------- Configuration ----------- #

class Configuration(object):
    def __init__(self):
        """
        Represents the default configuration which will be used by Archey.
        Values present in the dictionary below are needed.
        New optional values may be added with `_updateRecursive()` method.
        """
        self.config = {
            'colors_palette': {
                'use_unicode': False
            },
            'default_strings': {
                'no_address': 'No Address',
                'not_detected': 'Not detected',
                'virtual_environment': 'Virtual Environment',
                'bare_metal_environment': 'Bare-metal Environment'
            },
            'ip_settings': {
                'lan_ip_max_count': 2,
                'wan_ip_v6_support': True
            },
            'temperature': {
                'char_before_unit': ' ',
                'use_fahrenheit': False
            },
            'timeout': {
                'ipv4_detection': 1,
                'ipv6_detection': 1
            }
        }

        # Let's "save" `STDERR` file descriptor for `suppress_warnings` option
        self._stderr = sys.stderr

        # Now, let's load each optional configuration file in a "regular" order
        self.loadConfiguration('/etc/archey4/')
        self.loadConfiguration(os.path.expanduser('~') + '/.config/archey4/')
        self.loadConfiguration(os.path.dirname(os.path.realpath(__file__)))

    def get(self, key, default=None):
        """
        A binding method to imitate the `dict.get()` behavior.
        """
        return self.config.get(key, default)

    def loadConfiguration(self, path):
        # If a previous configuration file has denied overriding...
        if not self.config.get('allow_overriding', True):
            #  ... don't load this one.
            return

        if not path.endswith('/'):
            path += '/'

        path += 'config.json'

        try:
            with open(path) as file:
                self._updateRecursive(self.config, json.load(file))

            # If the user does not want any warning to appear : 2> /dev/null
            if self.config.get('suppress_warnings', False):
                # One more if statement to avoid multiple `open` calls.
                if sys.stderr == self._stderr:
                    sys.stderr = open(os.devnull, 'w')

            else:
                # One more if statement to avoid useless assignments and...
                # ... for closing previously opened new file descriptor.
                if sys.stderr != self._stderr:
                    sys.stderr.close()
                    sys.stderr = self._stderr

        except FileNotFoundError:
            pass

        # For backward compatibility with Python versions prior to 3.5.0
        #   we use `ValueError` instead of `json.JSONDecodeError`.
        except ValueError as e:
            print('Warning: {0} ({1})'.format(e, path), file=sys.stderr)

    def _updateRecursive(self, oldDict, newDict):
        """
        A method for recursively merging dictionaries as...
        ... `dict.update()` is not able to do this.
        Original snippet taken from here :
        https://gist.github.com/angstwad/bf22d1822c38a92ec0a9
        """
        for key, value in newDict.items():
            if key in oldDict and isinstance(oldDict[key], dict) \
                    and isinstance(value, dict):
                self._updateRecursive(oldDict[key], value)

            else:
                oldDict[key] = value


# ---------- Global variables --------- #

# We create a global instance of our `Configuration` Class
config = Configuration()

# We'll list the running processes only one time
try:
    PROCESSES = check_output(
        ['ps', '-u' + str(getuid()) if getuid() != 0 else '-ax',
         '-o', 'comm', '--no-headers'], universal_newlines=True
    ).splitlines()

except FileNotFoundError:
    print('Please, install first `procps` on your distribution.',
          file=sys.stderr)
    exit()


# ----------- Output handler ---------- #

class Output(object):
    def __init__(self):
        try:
            lsbOutput = check_output(
                ['lsb_release', '-i', '-s'],
                universal_newlines=True
            ).rstrip()

        except FileNotFoundError:
            print('Please, install first `lsb-release` on your distribution.',
                  file=sys.stderr)
            exit()

        if re.search(
               'Microsoft',
               check_output(['uname', '-r'], universal_newlines=True)):
            self.distribution = Distributions.WINDOWS

        else:
            for distribution in Distributions:
                if re.fullmatch(distribution.value, lsbOutput):
                    self.distribution = distribution
                    break

            else:
                self.distribution = Distributions.LINUX

        # Each class output will be added in the list below afterwards
        self.results = []

    def append(self, key, value):
        self.results.append(
            '{0}{1}:{2} {3}'.format(
                colorDict[self.distribution][1], key, colorDict['clear'], value
            )
        )

    def output(self):
        # Let's center the entries according to the logo (handles odd numbers)
        self.results[0:0] = [''] * ((18 - len(self.results)) // 2)
        self.results.extend([''] * (18 - len(self.results)))

        try:
            print(
                logosDict[self.distribution].format(
                    c=colorDict[self.distribution],
                    r=self.results
                ) + colorDict['clear']
            )

        except UnicodeError:
            print(
                'Your locale or TTY seems not supporting UTF8 encoding.\n'
                'Please disable Unicode within your configuration file.',
                file=sys.stderr
            )


# -------------- Entries -------------- #

class User:
    def __init__(self):
        self.value = getenv(
            'USER',
            config.get('default_strings')['not_detected']
        )


class Hostname:
    def __init__(self):
        self.value = check_output(
            ['uname', '-n'],
            universal_newlines=True
        ).rstrip()


class Model:
    def __init__(self):
        try:
            with open('/sys/devices/virtual/dmi/id/product_name') as file:
                model = file.read().rstrip()

        except FileNotFoundError:
            # The file above does not exist, is this device a Raspberry Pi ?
            # Let's retrieve the Hardware and Revision IDs with `/proc/cpuinfo`
            with open('/proc/cpuinfo') as file:
                output = file.read()

            hardware = re.search('(?<=Hardware\t: ).*', output)
            revision = re.search('(?<=Revision\t: ).*', output)

            # If the output contains 'Hardware' and 'Revision'...
            if hardware and revision:
                # ... let's set a pretty info string with these data
                model = 'Raspberry Pi {0} (Rev. {1})'.format(
                    hardware.group(0),
                    revision.group(0)
                )

            else:
                # A tricky way to retrieve some details about hypervisor...
                # ... within virtual contexts.
                # `archey` needs to be run as root although.
                try:
                    virtWhat = ', '.join(
                        check_output(
                            ['virt-what'],
                            stderr=DEVNULL, universal_newlines=True
                        ).splitlines()
                    )

                    if virtWhat:
                        try:
                            # Sometimes we may gather info added by...
                            # ... hosting service provider this way
                            model = check_output(
                                ['dmidecode', '-s', 'system-product-name'],
                                stderr=DEVNULL, universal_newlines=True
                            ).rstrip()

                        except (FileNotFoundError, CalledProcessError):
                            model = config.get(
                                'default_strings'
                            )['virtual_environment']

                        model += ' ({0})'.format(virtWhat)

                    else:
                        model = config.get(
                            'default_strings'
                        )['bare_metal_environment']

                except (FileNotFoundError, CalledProcessError):
                    model = config.get('default_strings')['not_detected']

        self.value = model


class Distro:
    def __init__(self):
        self.value = '{0} {1}'.format(
            check_output(
                ['lsb_release', '-d', '-s'],
                universal_newlines=True
            ).rstrip(),
            check_output(
                ['uname', '-m'],
                universal_newlines=True
            ).rstrip()
        )


class Kernel:
    def __init__(self):
        self.value = check_output(
            ['uname', '-r'],
            universal_newlines=True
        ).rstrip()


class Uptime:
    def __init__(self):
        with open('/proc/uptime') as file:
            fuptime = int(file.read().split('.')[0])

        day, fuptime = divmod(fuptime, 86400)
        hour, fuptime = divmod(fuptime, 3600)
        minute = fuptime // 60

        uptime = ''
        uptime += ((str(day) + ' day' + ('s' if day > 1 else '') + ('' if minute == 0 and hour == 0 else (' and ' if (minute != 0 and hour == 0) or (minute == 0 and hour != 0) else ', '))) if day >= 1 else '')
        uptime += ((str(hour) + ' hour' + ('s' if hour > 1 else '') + (' and ' if minute != 0 else '')) if hour >= 1 else '')
        uptime += ((str(minute) + ' minute' + ('s' if minute > 1 else '')) if minute >= 1 else ('< 1 minute' if day == 0 and hour == 0 else ''))

        self.value = uptime


class WindowManager:
    def __init__(self):
        try:
            wm = re.search(
                '(?<=Name: ).*',
                check_output(
                    ['wmctrl', '-m'],
                    stderr=DEVNULL, universal_newlines=True
                )
            ).group(0)

        except (FileNotFoundError, CalledProcessError):
            for key in wmDict.keys():
                if key in PROCESSES:
                    wm = wmDict[key]
                    break

            else:
                wm = config.get('default_strings')['not_detected']

        self.value = wm


class DesktopEnvironment:
    def __init__(self):
        for key in deDict.keys():
            if key in PROCESSES:
                de = deDict[key]
                break

        else:
            # Let's rely on an environment var if the loop above didn't `break`
            de = getenv(
                'XDG_CURRENT_DESKTOP',
                config.get('default_strings')['not_detected']
            )

        self.value = de


class Shell:
    def __init__(self):
        self.value = getenv(
            'SHELL',
            config.get('default_strings')['not_detected']
        )


class Terminal:
    def __init__(self):
        terminal = getenv(
            'TERM',
            config.get('default_strings')['not_detected']
        )

        # On systems with non-Unicode locales, we imitate '\u2588' character
        # ... with '#' to display the terminal colors palette.
        # This is the default option for backward compatibility.
        colors = ' '.join([
            '\x1b[0;3{0}m{1}\x1b[1;3{0}m{1}{2}'.format(
                i,
                '\u2588' if config.get('colors_palette')['use_unicode']
                else '#',
                colorDict['clear']
            ) for i in range(7, 0, -1)
        ])

        self.value = '{0} {1}'.format(terminal, colors)


class Temperature:
    def __init__(self):
        temps = []

        try:
            # Let's try to retrieve a value from the Broadcom chip on Raspberry
            temp = float(
                re.search(
                    '\d+\.\d+',
                    check_output(
                        ['/opt/vc/bin/vcgencmd', 'measure_temp'],
                        stderr=DEVNULL, universal_newlines=True
                    )
                ).group(0)
            )

            temps.append(
                self._convertToFahrenheit(temp)
                if config.get('temperature')['use_fahrenheit'] else temp
            )

        except (FileNotFoundError, CalledProcessError):
            pass

        # Now we just check for values within files present in the path below
        for thermalFile in glob('/sys/class/thermal/thermal_zone*/temp'):
            with open(thermalFile) as file:
                temp = float(file.read().strip()) / 1000
                if temp != 0.0:
                    temps.append(
                        self._convertToFahrenheit(temp)
                        if config.get('temperature')['use_fahrenheit']
                        else temp
                    )

        if temps:
            self.value = '{0}{2}{3} (Max. {1}{2}{3})'.format(
                str(round(sum(temps) / len(temps), 1)),
                str(round(max(temps), 1)),
                config.get('temperature')['char_before_unit'],
                'F' if config.get('temperature')['use_fahrenheit'] else 'C'
            )

        else:
            self.value = config.get('default_strings')['not_detected']

    """
    Simple Celsius to Fahrenheit conversion method
    """
    def _convertToFahrenheit(self, temp):
        return temp * (9 / 5) + 32


class Packages:
    def __init__(self):
        for packagesTool in [['dnf',    'list',   'installed'],
                             ['dpkg',   '--get-selections'],
                             ['emerge', '-ep',    'world'],
                             ['pacman', '-Q'],
                             ['rpm',    '-qa'],
                             ['yum',    'list',   'installed'],
                             ['zypper', 'search', '-i']]:
            try:
                results = check_output(
                    packagesTool,
                    stderr=DEVNULL, env={'LANG': 'C'}, universal_newlines=True
                )
                packages = results.count('\n')

                if 'dnf' in packagesTool:  # Deduct extra heading line
                    packages -= 1

                elif 'dpkg' in packagesTool:  # Packages removed but not purged
                    packages -= results.count('deinstall')

                elif 'emerge' in packagesTool:  # Deduct extra heading lines
                    packages -= 5

                elif 'yum' in packagesTool:  # Deduct extra heading lines
                    packages -= 2

                elif 'zypper' in packagesTool:  # Deduct extra heading lines
                    packages -= 5

                break

            except (FileNotFoundError, CalledProcessError):
                pass

        else:
            packages = config.get('default_strings')['not_detected']

        self.value = packages


class CPU:
    def __init__(self):
        model_name_regex = re.compile(
            '^model name\s*:\s*(.*)$',
            flags=re.IGNORECASE | re.MULTILINE
        )

        with open('/proc/cpuinfo') as file:
            cpuinfo = re.search(model_name_regex, file.read())

        # This test case has been built for some ARM architectures (see #29).
        # Sometimes, `model name` info is not present within `/proc/cpuinfo`.
        # We use the output of `lscpu` program (util-linux-ng) to retrieve it.
        if not cpuinfo:
            cpuinfo = re.search(
                model_name_regex,
                check_output(
                    ['lscpu'],
                    env={'LANG': 'C'}, universal_newlines=True
                )
            )

        # Sometimes CPU model name contains extra ugly white-spaces.
        self.value = re.sub('\s+', ' ', cpuinfo.group(1))


class GPU:
    def __init__(self):
        """
        Some explanations are needed here :
        * We call `lspci` program to retrieve hardware devices
        * We keep only the entries with "3D", "VGA" or "Display"
        * We sort them in the same order as above (for relevancy)
        """
        try:
            lspci_output = sorted([
                (i.split(': ')[0].split(' ')[1], i.split(': ')[1])
                for i in check_output(
                    ['lspci'], universal_newlines=True
                ).splitlines()
                if '3D' in i or 'VGA' in i or 'Display' in i
                ], key=lambda x: len(x[1])
            )

            if lspci_output:
                gpuinfo = lspci_output[0][1]

                # If the line got too long, let's truncate it and add some dots
                if len(gpuinfo) > 48:
                    # This call truncates `gpuinfo` with words preservation
                    gpuinfo = re.search(
                        '.{1,45}(?:\W|$)', gpuinfo
                    ).group(0).strip() + '...'

            else:
                gpuinfo = config.get('default_strings')['not_detected']

        except (FileNotFoundError, CalledProcessError):
            gpuinfo = config.get('default_strings')['not_detected']

        self.value = gpuinfo


class RAM:
    def __init__(self):
        try:
            ram = ''.join(
                filter(
                    re.compile('Mem').search,
                    check_output(
                        ['free', '-m'],
                        env={'LANG': 'C'}, universal_newlines=True
                    ).splitlines()
                )
            ).split()
            used = float(ram[2])
            total = float(ram[1])

        except (IndexError, FileNotFoundError):
            # An in-digest one-liner to retrieve memory info into a dictionary
            with open('/proc/meminfo') as file:
                ram = {
                    i.split(':')[0]: float(i.split(':')[1].strip(' kB')) / 1024
                    for i in filter(None, file.read().splitlines())
                }

            total = ram['MemTotal']
            # Here, let's imitate the `free` command behavior
            # (https://gitlab.com/procps-ng/procps/blob/master/proc/sysinfo.c#L787)
            used = total - (ram['MemFree'] + ram['Cached'] + ram['Buffers'])
            if used < 0:
                used += ram['Cached'] + ram['Buffers']

        self.value = '{0}{1} MB{2} / {3} MB'.format(
            colorDict['sensors'][int(((used / total) * 100) // 33.34)],
            int(used),
            colorDict['clear'],
            int(total)
        )


class Disk:
    def __init__(self):
        total = re.sub(
            ',', '.',
            check_output([
                'df', '-Tlh', '-B', 'GB', '--total',
                '-t', 'ext4', '-t', 'ext3', '-t', 'ext2',
                '-t', 'reiserfs', '-t', 'jfs', '-t', 'zfs',
                '-t', 'ntfs', '-t', 'fat32', '-t', 'btrfs',
                '-t', 'fuseblk', '-t', 'xfs',
                '-t', 'simfs', '-t', 'tmpfs'
                ], universal_newlines=True
            ).splitlines()[-1]
        ).split()

        self.value = '{0}{1}{2} / {3}'.format(
            colorDict['sensors'][int(float(total[5][:-1]) // 33.34)],
            re.sub('GB', ' GB', total[3]), colorDict['clear'],
            re.sub('GB', ' GB', total[2])
        )


class LAN_IP:
    def __init__(self):
        try:
            addresses = check_output(
                ['hostname', '-I'],
                stderr=DEVNULL, universal_newlines=True
            ).split()

        except (CalledProcessError, FileNotFoundError):
            # Slow manual workaround for old `inetutils` versions, with `ip`
            addresses = check_output(
                    ['cut', '-d', ' ', '-f', '4'],
                    universal_newlines=True,
                    stdin=Popen(['cut', '-d', '/', '-f', '1'],
                                stdout=PIPE,
                                stdin=Popen(['tr', '-s', ' '],
                                            stdout=PIPE,
                                            stdin=Popen(['grep', '-E',
                                                         'scope (global|site)'
                                                         ], stdout=PIPE,
                                                        stdin=Popen(['ip',
                                                                     '-o',
                                                                     'addr',
                                                                     'show',
                                                                     'up'],
                                                                    stdout=PIPE
                                                                    ).stdout
                                                        ).stdout
                                            ).stdout
                                ).stdout
            ).splitlines()

        # Use list slice to save only `lan_ip_max_count` from `addresses`.
        # If set to `False`, don't modify the list.
        # This option is still optional.
        self.value = ', '.join(
            addresses[:(
                config.get('ip_settings')['lan_ip_max_count']
                if config.get('ip_settings')['lan_ip_max_count'] is not False
                else len(addresses)
            )]
        ) or config.get('default_strings')['no_address']


class WAN_IP:
    def __init__(self):
        # IPv6 address retrieval (unless the user doesn't want it).
        if config.get('ip_settings')['wan_ip_v6_support']:
            try:
                ipv6_value = check_output([
                    'dig', '+short', '-6', 'aaaa', 'myip.opendns.com',
                    '@resolver1.ipv6-sandbox.opendns.com'
                    ], timeout=config.get('timeout')['ipv6_detection'],
                    stderr=DEVNULL, universal_newlines=True
                ).rstrip()

            except (FileNotFoundError, TimeoutExpired, CalledProcessError):
                try:
                    ipv6_value = check_output([
                        'wget', '-qO-', 'https://v6.ident.me/'
                        ], timeout=config.get('timeout')['ipv6_detection'],
                        universal_newlines=True
                    )

                except (CalledProcessError, TimeoutExpired):
                    # It looks like this user doesn't have any IPv6 address...
                    # ... or is not connected to Internet.
                    ipv6_value = None

                except FileNotFoundError:
                    ipv6_value = None
                    print('Warning: `wget` has not been found on your system.',
                          file=sys.stderr)

        else:
            ipv6_value = None

        # IPv4 addresses retrieval (anyway).
        try:
            ipv4_value = check_output([
                'dig', '+short', 'myip.opendns.com', '@resolver1.opendns.com'
                ], timeout=config.get('timeout')['ipv4_detection'],
                stderr=DEVNULL, universal_newlines=True
            ).rstrip()

        except (FileNotFoundError, TimeoutExpired, CalledProcessError):
            try:
                ipv4_value = check_output([
                    'wget', '-qO-', 'https://v4.ident.me/'
                    ], timeout=config.get('timeout')['ipv4_detection'],
                    universal_newlines=True
                )

            except (CalledProcessError, TimeoutExpired):
                # This user looks not connected to Internet...
                ipv4_value = None

            except FileNotFoundError:
                ipv4_value = None
                # If statement so as to not print the same message twice.
                if not config.get('ip_settings')['wan_ip_v6_support']:
                    print('Warning: `wget` has not been found on your system.',
                          file=sys.stderr)

        self.value = ', '.join(
            filter(None, (ipv4_value, ipv6_value))
        ) or config.get('default_strings')['no_address']


# ----------- Classes Index ----------- #

class Classes(Enum):
    User = User
    Hostname = Hostname
    Model = Model
    Distro = Distro
    Kernel = Kernel
    Uptime = Uptime
    WindowManager = WindowManager
    DesktopEnvironment = DesktopEnvironment
    Shell = Shell
    Terminal = Terminal
    Packages = Packages
    Temperature = Temperature
    CPU = CPU
    GPU = GPU
    RAM = RAM
    Disk = Disk
    LAN_IP = LAN_IP
    WAN_IP = WAN_IP


# ---------------- Main --------------- #

def main():
    output = Output()
    for key in Classes:
        if config.get('entries', {}).get(key.name, True):
            output.append(key.name, key.value().value)

    output.output()


if __name__ == '__main__':
    main()
