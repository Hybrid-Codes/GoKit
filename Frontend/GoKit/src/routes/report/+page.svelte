<script>
    let pdfUrl = '';

    async function generatePDF() {
      const response = await fetch('http://localhost:8080/generate-pdf', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/pdf'
        }
      });

      if (response.ok) {
        const blob = await response.blob();
        pdfUrl = window.URL.createObjectURL(blob); // Create URL to show the PDF in iframe
      } else {
        console.error("Failed to generate PDF");
      }
    }
  </script>

  {#if pdfUrl}
    <iframe src={pdfUrl} width="100%" height="600px"></iframe> <!-- Displays the PDF -->
  {/if}

  <button on:click={generatePDF}
    class="w-full px-4 py-2 bg-blue-600
     text-white font-semibold rounded-md shadow-sm
     hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2
     focus:ring-blue-500 mt-7">Generate PDF</button>
