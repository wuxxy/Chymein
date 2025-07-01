<script lang="ts">

    import {theme} from "$lib/stores/theme";

    interface ButtonProps {
        label: string;
        onClick: () => void;
        disabled?: boolean;
        style: string
    }

    let props: ButtonProps = $props();

</script>

<button
        onclick={props.onClick}
        class="px-4 py-2 dark:text-white font-semibold bg-white dark:bg-black w-min"
        disabled={props.disabled}
        style={props.style}
>
    {props.label}
</button>
