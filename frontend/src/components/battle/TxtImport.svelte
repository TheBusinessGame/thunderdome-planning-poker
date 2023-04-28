<script>

    import HollowButton from '../HollowButton.svelte'
    import { AppConfig } from '../../config.js'
    import { _ } from '../../i18n.js'

    export let notifications
    export let eventTag = () => {}
    export let handlePlanAdd = () => {}

    const allowTxtImport = AppConfig.AllowTxtImport

    let plans = []

    function uploadFile() {
        let file = this.files[0]
        if (!file) {
            return
        }
        if (file.type !== 'text/plain') {
            notifications.danger($_('importTxtFileBadFileTypeError'))
            eventTag('Txt_import_failed', 'battle', `file.type not text/plain`)
            return
        }

        let reader = new FileReader()

        reader.readAsText(file)

        reader.onload = () => {
            try {
                const content = reader.result;
                const items = parseTextFile(content);
                if (items) {
                    const totalItems = items.length
                    for (let i = 0; i < totalItems; i++) {
                        const item = items[i]
                        const plan = extractPlanData(item);
                        plans.push(plan)
                        handlePlanAdd(plan);
                    }
                    eventTag(
                        'Txt_import_success',
                        'battle',
                        `total stories imported: ${totalItems}`,
                    )
                }
            } catch (e) {
                notifications.danger($_('importTxtFileReadFileError'))
                eventTag('Txt_import_failed', 'battle', `ferror reading file`)
            }
        }

        reader.onerror = () => {
            notifications.danger($_('importTxtFileReadFileError'))
            eventTag('Txt_import_failed', 'battle', `ferror reading file`)
        }
    }

    function parseTextFile(content) {
        const lines = content.split('\n');
        const items = [];

        for (let i = 0; i < lines.length; i++) {
            const line = lines[i].trim();
            if (line) {
                items.push(line);
            }
        }

        return items;
    }

    function extractPlanData(item) {
        const fields = item.split(',');
        const plan = {
            id: '',
            planName: fields[0].trim(),
            type: fields[1].trim().toLowerCase(),
            referenceId: fields[2].trim(),
            link: fields[3].trim(),
            description: fields[4].trim(),
            acceptanceCriteria: fields[5] ? fields[5].trim() : ''
        };

        return plan;
    }
</script>

{#if allowTxtImport}
    <HollowButton
        type="label"
        additionalClasses="rtl:ml-2 ltr:mr-2"
        color="purple"
        labelFor="txtimport"
    >
        {$_('importTxt')}
        <input
            type="file"
            on:change="{uploadFile}"
            class="hidden"
            id="txtimport"
            accept=".txt"
        />
    </HollowButton>
{/if}
