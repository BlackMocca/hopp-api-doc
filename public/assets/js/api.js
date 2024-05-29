async function getTeamCollection() {
    try {
        const response = await fetch('/team/collections', {
          method: 'GET',
        });
  
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
  
        return await response.text();
      } catch (error) {
        console.error('Error at fetch user collection:', error);
        // Handle errors
    }
}

async function getUserCollection() {
    try {
        const response = await fetch('/my/collection', {
          method: 'GET',
        });
  
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
  
        return await response.text();
      } catch (error) {
        console.error('Error at fetch user collection:', error);
        // Handle errors
    }
}

async function signout() {
    try {
        const response = await fetch('/signout', {
          method: 'POST',
        });
  
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
  
        return response.status
      } catch (error) {
        console.error('Error at fetch user collection:', error);
        // Handle errors
    }
}

function authProvider(provider) {
  const searchParams = new URLSearchParams();
  searchParams.set("provider", provider)
  window.location.href = "/v1/auth/signin?" + searchParams.toString()
}

async function importCollection(file) {
  const formData = new FormData()
  formData.append("file", file)

  try {
    const response = await fetch("/import", {
      method: 'POST',
      body: formData,
    })
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    return await response.text();
  } catch (error) {
    console.error('Error at import collection:', error);
  }
}

async function apiDeleteMyCollection(collectionId) {
  try {
      const response = await fetch('/my/collection/'+collectionId, {
        method: 'DELETE',
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      return await response.text();
    } catch (error) {
      console.error('Error at delete my collection:', error);
      // Handle errors
  }
}