package services

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
	"google.golang.org/api/option"
)

// TranscribeBase64Audio toma un audio codificado en Base64, lo decodifica y lo envía a la API Speech-to-Text para su transcripción en inglés.
func TranscribeBase64Audio(base64Audio string) string {

	// Establecer la variable de entorno para las credenciales
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./credenciales/my-project-go-423503-373c395b59da.json")

	// Decodificar el audio Base64 a bytes
	data, err := base64.StdEncoding.DecodeString(base64Audio)
	if err != nil {
		log.Fatalf("Failed to decode base64 audio: %v", err)
	}

	// Guardar el archivo de audio decodificado (opcional, para verificación)
	err = ioutil.WriteFile("./output_decoded.wav", data, 0644)
	if err != nil {
		log.Println("Error al guardar el archivo de audio decodificado:", err)
	}

	// Verificar el contenido del archivo decodificado
	audioFile := "./output_decoded.wav"
	audioFileBytes, err := ioutil.ReadFile(audioFile)
	if err != nil {
		log.Println("Error al leer el archivo de audio:", err)
	}

	// Configurar contexto y cliente de la API de Google Cloud Speech-to-Text
	ctx := context.Background()
	client, err := speech.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		log.Println("Error al crear el cliente de la API:", err)
	}

	// Configurar la solicitud de reconocimiento para inglés
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:                   speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz:            24000,
			LanguageCode:               "en-US", // Cambio aquí para reconocimiento en inglés
			MaxAlternatives:            1,       // Obtener solo una alternativa de resultado
			EnableAutomaticPunctuation: true,    // Habilitar puntuación automática
			EnableWordTimeOffsets:      true,    // Habilitar tiempos de palabras
			Model:                      "default",
			UseEnhanced:                true,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{
				Content: audioFileBytes,
			},
		},
	}

	// Realizar la solicitud de reconocimiento
	resp, err := client.Recognize(ctx, req)
	if err != nil {
		log.Fatalf("Failed to recognize: %v", err)
	}

	// Procesar los resultados de la transcripción
	if len(resp.Results) > 0 {
		var transcribedText string
		for _, result := range resp.Results {
			for _, alt := range result.Alternatives {
				transcribedText += alt.Transcript
			}
		}
		return transcribedText
	}

	return ""
}
