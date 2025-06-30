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