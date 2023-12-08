using NUnit.Framework;

namespace pe_homework_7_tests;

public class Tests
{
    [Test]
    public void TestPherma1()
    {
        int x = 2;
        int y = 3;
        int z = 4;

        int x3 = (int)Math.Pow(x, 3);
        int y3 = (int)Math.Pow(y, 3);
        int z3 = (int)Math.Pow(z, 3);
        
        Assert.That(x3 + y3 != z3);
    }
    
    [Test]
    public void TestPherma2()
    {
        int x = 1;
        int y = 1;
        int z = 1;

        int x3 = (int)Math.Pow(x, 3);
        int y3 = (int)Math.Pow(y, 3);
        int z3 = (int)Math.Pow(z, 3);
        
        Assert.That(x3 + y3 != z3);
    }
    
    [Test]
    public void TestPherma3()
    {
        int x = 5;
        int y = 5;
        int z = 6;

        int x3 = (int)Math.Pow(x, 4);
        int y3 = (int)Math.Pow(y, 4);
        int z3 = (int)Math.Pow(z, 4);
        
        Assert.That(x3 + y3 != z3);
    }
}