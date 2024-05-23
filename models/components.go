package models

const IndexForm = `
{{ define "content" }}
<div class="w-full max-w-2xl lg:max-w-3xl" id="indexDiv">
    <form autocomplete="off" hx-post="/shorten/" hx-target="#indexDiv" hx-swap="outerHTML">
        <div class="innerForm mx-auto mt-3 relative min-w-0 max-w-8xl flex items-center justify-center py-2 px-2 rounded-2xl gap-2">
            <input type="text" name="url" placeholder="Paste your long URL here" class="focus:outline-none px-6 py-3 rounded-md min-w-[3rem] flex-grow" required>
            <button type="submit" class="bbtn focus:outline-none rounded-xl px-8 py-3">Generate</button>
        </div>
    </form>
</div>
{{ end }}
`

const ResultForm = `
<div class="w-full max-w-2xl lg:max-w-3xl text-gray-500" id="resultDiv">
    <form autocomplete="off" hx-get="/" hx-target="#resultDiv" hx-swap="outerHTML">

        <div class="innerForm mx-auto mt-3 relative min-w-0 max-w-8xl flex items-center justify-center py-2 px-2 rounded-2xl gap-2">
            <button onclick="copyToClipboard()" class="pl-3 px-1 py-3 bg-transparent transition-colors duration-300 hover:text-gray-200 hover:border-gray-200 " type="button">
            <span id="default-icon">
                <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 20">
                    <path d="M16 1h-3.278A1.992 1.992 0 0 0 11 0H7a1.993 1.993 0 0 0-1.722 1H2a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2Zm-3 14H5a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2Zm0-4H5a1 1 0 0 1 0-2h8a1 1 0 1 1 0 2Zm0-5H5a1 1 0 0 1 0-2h2V2h4v2h2a1 1 0 1 1 0 2Z"/>
                </svg>
            </span>
            <span id="success-icon" class="hidden ">
                <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 12">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5.917 5.724 10.5 15 1.5"/>
                </svg>
            </span>
            </button>
            <p type="text" id="urlOut" class="focus:outline-none px-1 py-3 text-gray-400">%s</p>
            <a href="%s" target="_blank"><button id="visit" class="focus:outline-none rounded-xl px-2 py-3 transition-colors duration-300 hover:text-gray-200 hover:border-gray-200">Visit</button></a>
            <button type="submit" class="bbtn focus:outline-none rounded-xl px-4 py-3 md:px-8">Generate</button>
        </div>
    </form>
</div>
`
const IndexHX = `
<div class="w-full max-w-2xl lg:max-w-3xl" id="indexDiv">
    <form autocomplete="off" hx-post="/shorten/" hx-target="#indexDiv" hx-swap="outerHTML">
        <div class="innerForm mx-auto mt-3 relative min-w-0 max-w-8xl flex items-center justify-center py-2 px-2 rounded-2xl gap-2">
            <input type="text" name="url" placeholder="Paste your long URL here" class="focus:outline-none px-6 py-3 rounded-md min-w-[3rem] flex-grow" required>
            <button type="submit" class="bbtn focus:outline-none rounded-xl px-8 py-3">Generate</button>
        </div>
    </form>
</div>
`
