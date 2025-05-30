document.getElementById('convertBtn').addEventListener('click', async () => {
    const amount = parseFloat(document.getElementById('amount').value);
    const fromCurrency = document.getElementById('fromCurrency').value;
    const toCurrency = document.getElementById('toCurrency').value;

    if (isNaN(amount)) {
        alert("Digite um valor válido!");
        return;
    }

    try {
        // Chama sua API local
        const response = await fetch(`http://localhost:8080/api/rates/${fromCurrency}`);
        const data = await response.json();

        if (!data.rates || !data.rates[toCurrency]) {
            throw new Error("Moeda não encontrada");
        }

        const rate = data.rates[toCurrency];
        const convertedValue = (amount * rate).toFixed(2);

        // Exibe o resultado
        document.getElementById('convertedValue').textContent = 
            `${amount} ${fromCurrency} = ${convertedValue} ${toCurrency}`;
    } catch (error) {
        console.error("Erro na conversão:", error);
        alert("Erro ao converter moeda. Verifique o console para detalhes.");
    }
});

document.getElementById('convertBtn').addEventListener('click', async () => {
            const amount = parseFloat(document.getElementById('amount').value);
            const fromCurrency = document.getElementById('fromCurrency').value;
            const toCurrency = document.getElementById('toCurrency').value;

            if (isNaN(amount) || amount <= 0) {
                alert("Digite um valor válido (maior que zero)!");
                return;
            }

            try {
                const response = await fetch(`http://localhost:8080/api/rates/${fromCurrency}`);
                
                if (!response.ok) {
                    throw new Error(`Erro HTTP: ${response.status}`);
                }

                const data = await response.json();
                
                if (!data.rates || !data.rates[toCurrency]) {
                    throw new Error(`Moeda ${toCurrency} não encontrada`);
                }

                const rate = data.rates[toCurrency];
                const convertedValue = (amount * rate).toFixed(2);
                document.getElementById('convertedValue').textContent = 
                    `${amount} ${fromCurrency} = ${convertedValue} ${toCurrency}`;
            } catch (error) {
                console.error("Erro detalhado:", error);
                alert(`Erro ao converter: ${error.message}`);
            }
        });