import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { MessageSquare, FileText, ArrowRight, Brain, BookOpen, Target } from "lucide-react";

export default function LandingPage() {
  return (
    <div className="min-h-screen bg-white">
      <header className="absolute top-0 w-full z-50 bg-white/80 backdrop-blur-md border-b border-gray-100">
        <div className="container mx-auto py-3 px-6 flex justify-between items-center">
          <div className="flex-1 flex items-center justify-start">
            <Link
              href="/"
              className="text-[#E12D39] hover:text-[#c0252f] text-xl font-bold transition-colors"
            >
              poznawca
            </Link>
          </div>
          <div className="flex items-center gap-4">
            <Link href="/login">
              <Button variant="ghost" className="text-[#0C3B5F] hover:text-[#E12D39] font-medium">
                Sign In
              </Button>
            </Link>
            <Link href="/register">
              <Button className="bg-[#E12D39] hover:bg-[#c0252f] rounded-full px-6 shadow-lg hover:shadow-xl transition-all">
                Get Started
              </Button>
            </Link>
          </div>
        </div>
      </header>

      <section className="pt-24 pb-20 px-6 relative overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-br from-[#E12D39]/5 via-white to-[#0C3B5F]/5"></div>
        <div className="absolute top-20 right-10 w-72 h-72 bg-[#E12D39]/10 rounded-full blur-3xl"></div>
        <div className="absolute bottom-20 left-10 w-96 h-96 bg-[#0C3B5F]/10 rounded-full blur-3xl"></div>

        <div className="container mx-auto text-center max-w-5xl relative z-10">
          <div className="mb-8">
            <div className="inline-flex items-center gap-2 bg-[#E12D39]/10 text-[#E12D39] px-4 py-2 rounded-full text-sm font-medium mb-4">
              <Brain className="h-4 w-4" />
              Knowledge-Based Learning Platform
            </div>

            <h1 className="text-6xl md:text-7xl font-bold text-[#0C3B5F] mb-8 leading-tight tracking-tight">
              Master Polish
              <span className="text-[#E12D39] block">Residency Knowledge</span>
            </h1>

            <p className="text-xl text-gray-600 mb-12 max-w-3xl mx-auto leading-relaxed font-light">
              Build comprehensive knowledge for your Karta Stalego Pobytu through interactive
              learning. Practice interviews and test your understanding of Polish culture, history,
              and society.
            </p>
          </div>

          <div className="flex flex-col sm:flex-row gap-6 justify-center mb-16">
            <Link href="/portal">
              <Button className="bg-[#E12D39] hover:bg-[#c0252f] text-white text-lg rounded-full shadow-xl hover:shadow-2xl transition-all hover:scale-105">
                Start Learning
                <ArrowRight className="ml-2 h-5 w-5" />
              </Button>
            </Link>
            <Link href="/">
              <Button
                variant="outline"
                className="border-2 border-[#0C3B5F]/20 text-[#0C3B5F] hover:bg-[#0C3B5F] hover:text-white text-lg rounded-full backdrop-blur-sm"
              >
                Explore Platform
              </Button>
            </Link>
          </div>

          {/* Knowledge Areas */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-4xl mx-auto">
            <div className="bg-white/60 backdrop-blur-sm border border-gray-100 rounded-2xl p-6 hover:bg-white/80 transition-all">
              <div className="bg-[#E12D39]/10 w-12 h-12 rounded-xl flex items-center justify-center mb-4 mx-auto">
                <BookOpen className="h-6 w-6 text-[#E12D39]" />
              </div>
              <h3 className="font-semibold text-[#0C3B5F] mb-2">Polish History</h3>
              <p className="text-gray-600 text-sm">Learn key historical events and figures</p>
            </div>

            <div className="bg-white/60 backdrop-blur-sm border border-gray-100 rounded-2xl p-6 hover:bg-white/80 transition-all">
              <div className="bg-[#0C3B5F]/10 w-12 h-12 rounded-xl flex items-center justify-center mb-4 mx-auto">
                <Target className="h-6 w-6 text-[#0C3B5F]" />
              </div>
              <h3 className="font-semibold text-[#0C3B5F] mb-2">Culture & Society</h3>
              <p className="text-gray-600 text-sm">Understand Polish traditions and values</p>
            </div>

            <div className="bg-white/60 backdrop-blur-sm border border-gray-100 rounded-2xl p-6 hover:bg-white/80 transition-all">
              <div className="bg-[#E12D39]/10 w-12 h-12 rounded-xl flex items-center justify-center mb-4 mx-auto">
                <Brain className="h-6 w-6 text-[#E12D39]" />
              </div>
              <h3 className="font-semibold text-[#0C3B5F] mb-2">Interview Skills</h3>
              <p className="text-gray-600 text-sm">Practice with realistic scenarios</p>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-24 px-6 bg-gray-50/50">
        <div className="container mx-auto max-w-7xl">
          <div className="text-center mb-20">
            <h2 className="text-5xl font-bold text-[#0C3B5F] mb-6 tracking-tight">
              Learn Through Practice
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto font-light">
              Our platform combines interactive learning with practical application to help you
              build the knowledge and confidence needed for your residency interview.
            </p>
          </div>

          <div className="grid lg:grid-cols-2 gap-12 items-center">
            {/* Mock Interviews Feature */}
            <div className="order-2 lg:order-1">
              <Card className="rounded-3xl border-0 shadow-2xl hover:shadow-3xl transition-all duration-500 overflow-hidden bg-white">
                <div className="bg-gradient-to-br from-[#E12D39]/5 to-[#E12D39]/10 p-10">
                  <div className="bg-[#E12D39] w-16 h-16 rounded-2xl flex items-center justify-center mb-8 shadow-lg">
                    <MessageSquare className="h-8 w-8 text-white" />
                  </div>
                  <h3 className="text-3xl font-bold text-[#0C3B5F] mb-6">
                    Interactive Mock Interviews
                  </h3>
                  <p className="text-gray-600 mb-8 leading-relaxed text-lg">
                    Practice with AI-powered interview simulations that adapt to your responses.
                    Build confidence through realistic conversations and receive detailed feedback
                    on your performance.
                  </p>
                  <div className="space-y-4">
                    <div className="flex items-center gap-4">
                      <div className="w-2 h-2 bg-[#E12D39] rounded-full"></div>
                      <span className="text-gray-700">
                        Multiple interviewer personalities and styles
                      </span>
                    </div>
                    <div className="flex items-center gap-4">
                      <div className="w-2 h-2 bg-[#E12D39] rounded-full"></div>
                      <span className="text-gray-700">Real-time language and content feedback</span>
                    </div>
                    <div className="flex items-center gap-4">
                      <div className="w-2 h-2 bg-[#E12D39] rounded-full"></div>
                      <span className="text-gray-700">Comprehensive performance analysis</span>
                    </div>
                  </div>
                </div>
              </Card>
            </div>

            {/* Knowledge Tests Feature */}
            <div className="order-1 lg:order-2">
              <Card className="rounded-3xl border-0 shadow-2xl hover:shadow-3xl transition-all duration-500 overflow-hidden bg-white">
                <div className="bg-gradient-to-br from-[#0C3B5F]/5 to-[#0C3B5F]/10 p-10">
                  <div className="bg-[#0C3B5F] w-16 h-16 rounded-2xl flex items-center justify-center mb-8 shadow-lg">
                    <FileText className="h-8 w-8 text-white" />
                  </div>
                  <h3 className="text-3xl font-bold text-[#0C3B5F] mb-6">
                    Comprehensive Knowledge Tests
                  </h3>
                  <p className="text-gray-600 mb-8 leading-relaxed text-lg">
                    Master Polish history, culture, and society through carefully crafted questions.
                    Track your learning progress and identify areas that need more attention.
                  </p>
                  <div className="space-y-4">
                    <div className="flex items-center gap-4">
                      <div className="w-2 h-2 bg-[#0C3B5F] rounded-full"></div>
                      <span className="text-gray-700">
                        Extensive question database covering all topics
                      </span>
                    </div>
                    <div className="flex items-center gap-4">
                      <div className="w-2 h-2 bg-[#0C3B5F] rounded-full"></div>
                      <span className="text-gray-700">
                        Adaptive difficulty based on your progress
                      </span>
                    </div>
                    <div className="flex items-center gap-4">
                      <div className="w-2 h-2 bg-[#0C3B5F] rounded-full"></div>
                      <span className="text-gray-700">Detailed explanations for every answer</span>
                    </div>
                  </div>
                </div>
              </Card>
            </div>
          </div>
        </div>
      </section>

      {/* Learning Approach Section */}
      <section className="py-24 px-6 bg-white">
        <div className="container mx-auto max-w-6xl">
          <div className="text-center mb-20">
            <h2 className="text-5xl font-bold text-[#0C3B5F] mb-6 tracking-tight">
              Knowledge-First Approach
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto font-light">
              We believe that true confidence comes from deep understanding. Our platform is
              designed to help you build comprehensive knowledge that will serve you well beyond the
              interview.
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            <div className="text-center group">
              <div className="bg-gradient-to-br from-[#E12D39]/10 to-[#E12D39]/5 w-20 h-20 rounded-2xl flex items-center justify-center mb-6 mx-auto group-hover:scale-110 transition-transform duration-300">
                <BookOpen className="h-10 w-10 text-[#E12D39]" />
              </div>
              <h3 className="text-xl font-bold text-[#0C3B5F] mb-4">Learn Fundamentals</h3>
              <p className="text-gray-600 leading-relaxed">
                Start with core concepts about Polish history, culture, and society. Build a solid
                foundation of knowledge.
              </p>
            </div>

            <div className="text-center group">
              <div className="bg-gradient-to-br from-[#0C3B5F]/10 to-[#0C3B5F]/5 w-20 h-20 rounded-2xl flex items-center justify-center mb-6 mx-auto group-hover:scale-110 transition-transform duration-300">
                <Target className="h-10 w-10 text-[#0C3B5F]" />
              </div>
              <h3 className="text-xl font-bold text-[#0C3B5F] mb-4">Practice Application</h3>
              <p className="text-gray-600 leading-relaxed">
                Apply your knowledge through interactive tests and mock interviews. See how concepts
                connect in real scenarios.
              </p>
            </div>

            <div className="text-center group">
              <div className="bg-gradient-to-br from-[#E12D39]/10 to-[#0C3B5F]/10 w-20 h-20 rounded-2xl flex items-center justify-center mb-6 mx-auto group-hover:scale-110 transition-transform duration-300">
                <Brain className="h-10 w-10 text-[#E12D39]" />
              </div>
              <h3 className="text-xl font-bold text-[#0C3B5F] mb-4">Build Confidence</h3>
              <p className="text-gray-600 leading-relaxed">
                Gain the confidence that comes from truly understanding the material, not just
                memorizing answers.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-white py-16 px-6 border-t border-gray-100">
        <div className="container mx-auto text-center">
          <div className="text-3xl font-bold text-[#E12D39] mb-6">poznawca</div>
          <p className="text-gray-600 mb-8 max-w-2xl mx-auto">
            Empowering learners with comprehensive knowledge for Polish residency success
          </p>
          <div className="flex justify-center gap-8 text-sm text-gray-500">
            <Link href="/privacy" className="hover:text-[#E12D39] transition-colors">
              Privacy Policy
            </Link>
            <Link href="/terms" className="hover:text-[#E12D39] transition-colors">
              Terms of Service
            </Link>
            <Link href="/contact" className="hover:text-[#E12D39] transition-colors">
              Contact Us
            </Link>
          </div>
        </div>
      </footer>
    </div>
  );
}
