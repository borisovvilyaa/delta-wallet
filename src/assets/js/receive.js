import qrcode from 'qrcode';
import WalletData from '../../../keys/wallet.json';

const address = WalletData.address; // Здесь укажите свой адрес
const qrCodeDiv = document.getElementById('qrcode'); // Получаем элемент, куда будем вставлять QR-код

qrcode.toDataURL(address, (err, url) => {
  if (err) {
    console.error(err);
    return;
  }

  const qrImg = document.createElement('img'); // Создаем элемент изображения
  qrImg.src = url; // Устанавливаем Data URL как src изображения
  qrImg.alt = 'QR Code'; // Устанавливаем атрибут alt для изображения
  qrCodeDiv.appendChild(qrImg); // Добавляем изображение в элемент #qrcode
});
