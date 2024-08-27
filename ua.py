import sys
import faker

def generate_user_agents(save_file, amount):
    fake = faker.Faker()
    user_agents = [fake.user_agent() for _ in range(amount)]
    
    with open(save_file, 'w') as f:
        f.write('\n'.join(user_agents))
    
    print(f"{amount} user agents saved to {save_file}")

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python ua.py [save file] [amount]")
        sys.exit(1)

    save_file = sys.argv[1]
    amount = int(sys.argv[2])

    generate_user_agents(save_file, amount)
