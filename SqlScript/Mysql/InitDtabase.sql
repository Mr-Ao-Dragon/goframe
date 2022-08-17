create table if not exists users
(
    Name            varchar(16),
    HexID           binary,
    UserMD5         binary,
    UserSHA1        binary,
    RegisterTime    datetime,
    LastLoginTime   datetime,
    BanDate         datetime,
    PermissionsID   binary,
    SaltedPassword  varchar(32),
    ClientPublicKey varchar(68),
    PRIMARY KEY (HexID)
);
create table if not exists Permissions
(
    PRIMARY KEY (PremissionsID),
    PremissionsID binary,
    Login         bool,
    Modshelf      bool,
    Modotheruser  bool,
    Modsetting    bool,
    IsAdmin       bool,
    IsSuperAdmin  bool,
    IsSystemUser  bool,
    BanUser       bool,
    UnbanUser     bool,
    BanIP         bool,
    UnbanIP       bool,
    GrantAdmin    bool
);
