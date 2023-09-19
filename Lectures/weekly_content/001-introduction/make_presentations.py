import os
import fnmatch
import subprocess

def remove_ext(filename):

    # Use os.path.splitext to split the filename and extension
    filename_without_extension, file_extension = os.path.splitext(filename)

    # Print the filename without the extension
    return filename_without_extension

# Specify the directory path you want to search in
directory_path = '.'

# Define the file extension you want to search for (e.g., '.md' for Markdown files)
file_extension = '.md'

# Use list comprehension to get a list of Markdown files in the directory
markdown_files = [f for f in os.listdir(directory_path) if fnmatch.fnmatch(f, f'*{file_extension}')]

# Print the list of Markdown files
for file in markdown_files:
    print(file)
    cleaned = remove_ext(file)
    # marp --theme gradient.css 02_user_input.md 02_user_input.pptx

    # Execute the 'marp' command to convert the Markdown file to pptx
    try:
        subprocess.run(['marp', '--theme gradient.css', '--pptx', file,cleaned+".pptx"])
        print(f'Successfully converted {file} to pptx.')
    except subprocess.CalledProcessError:
        print('Conversion failed. Please check if marp-cli is correctly installed.')

    try:
        subprocess.run(['marp', '--theme gradient.css', '--pdf', file,cleaned+".pdf"])
        print(f'Successfully converted {file} to pdf.')
    except subprocess.CalledProcessError:
        print('Conversion failed. Please check if marp-cli is correctly installed.')
