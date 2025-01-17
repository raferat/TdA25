<script lang="ts">
    import { page } from "$app/state";
</script>
<main style="padding: 50px;">
<center><h1>Při používání naší stránky jsme narazili na chybu</h1></center>
<h3>Jestli si myslíte, že je tato chyba způsobena námi, pošlete na naši podporu tento text:</h3>
<code>{page.url.searchParams.get("obj")}</code>

<br>
<center>
    <br>
    <br>
    <h2>Nastala chyba</h2><br>
    <img src="/chyba.png" alt="nastala chyba">
</center>

</main>