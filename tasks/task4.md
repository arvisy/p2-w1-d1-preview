Buffered Channel:
	~ produce goroutine tidak akan terblokir hingga 5 elemen pertama dikirim ke dalam channel.
	~ consume goroutine dapat mulai membaca dan mencetak nilai segera setelah elemen-elemen pertama dikirim.
	~ produce dan consume dapat terjadi secara asinkron, karena buffer memungkinkan beberapa elemen dikirim sebelum consume dimulai.
	~ program utama dapat keluar bahkan sebelum semua nilai dikonsumsi jika buffer tidak kosong.

Unbuffered Channel:
	~ produce akan terblokir hingga consume siap menerima setiap nilai yang dikirim.
	~ consume akan terblokir sampai ada nilai yang dikirim oleh produce.
	~ sifatnya lebih sinkron karena setiap elemen yang dikirim segera diterima oleh penerima.
	~ program utama akan menunggu hingga semua nilai dikonsumsi sebelum keluar.
