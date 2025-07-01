export type StatusType = {
    database: boolean;
    is_setup: boolean;
    port: string;
    time_alive: string;
    connected?:boolean
};

export type NavBarType = {
    label: string;
    href: string
}[]
export type UserType = {
    CreatedAt: string; // ISO timestamp
    UpdatedAt: string;
    DeletedAt: string | null;

    ID: string;
    Username: string;
    Email: string;
    Password: string;

    Permissions: Record<string, any> | null;

    Admin: boolean;
    SuperAdmin: boolean;

    LastActive: string;
    LoginAttempts: number;
    IsActive: boolean;
    IsVerified: boolean;
    TwoFactorEnabled: boolean;

    Sessions: any[] | null; // refine if session shape known
    Locked: boolean;
    DateOfBirth: string | null;

    Gender: 'unselected' | string;
    AvatarURL: string;
    BannerColor: string;
    Bio: string;
    PersonalLink: string;
    Signature: string;
    Language: string;
    Theme: string;

    Banned: boolean;
    Muted: boolean;
    AdminNotes: string;

    Metadata: Record<string, any> | null;
    PluginData: Record<string, any> | null;
};
