import ClipboardJS from 'clipboard';

document.addEventListener('DOMContentLoaded', function() {
  new ClipboardJS('[data-copy]', {
    text: function(trigger) {
      return trigger.getAttribute('data-copy');
    }
  }).on('success', function(e) {
    e.clearSelection();

    const notification = document.createElement('div');
    notification.classList.add('custom-notification', 'animate__animated', 'animate__fadeInRight');
    notification.innerHTML = 'Copied';

    document.body.appendChild(notification);

    setTimeout(function() {
      notification.classList.remove('animate__fadeInRight');
      notification.classList.add('animate__fadeOutRight');

      notification.addEventListener('animationend', function() {
        notification.remove();
      });
    }, 1000);
  });
});
