mport styles from '../styles/Gallery.module.css';

const Gallery = () => {
  // Replace with actual image data later
  const images = [
    { id: 1, src: '/images/image1.jpg', alt: 'Landscape 1' },
    { id: 2, src: '/images/image2.jpg', alt: 'Landscape 2' },
    { id: 3, src: '/images/image3.jpg', alt: 'Landscape 3' },
  ];

  return (
    <div className={styles.gallery}>
      {images.map((image) => (
        <img key={image.id} src={image.src} alt={image.alt} className={styles.galleryImage} />
      ))}
    </div>
  );
};

export default Gallery;
